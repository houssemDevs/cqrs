package cqrs

import (
	"fmt"
	"reflect"
	"sync"
)

type CommandDispatcher interface {
	RegisterHandler(handler CommandHandler, c Command, mds ...CommandMiddleware)
	Dispatch(c Command) error
	UseMiddleware(mds ...CommandMiddleware)
}

type inProcessCommandDispatcher struct {
	handlers map[string]CommandHandler
	middlewares []CommandMiddleware
}

func (cd *inProcessCommandDispatcher) RegisterHandler(handler CommandHandler, c Command, mds ...CommandMiddleware) {
	for _, md := range mds {
		handler = md(handler)
	}

	if registeredHandler, found := cd.handlers[c.CommandType()]; found {
		panic(fmt.Sprintf("A handler '%s' is already registred for command '%s'", reflect.TypeOf(registeredHandler), c.CommandType()))
	}

	cd.handlers[c.CommandType()] = handler
}

func (cd *inProcessCommandDispatcher) Dispatch(c Command) error {
	if handler, found := cd.handlers[c.CommandType()]; found {
		for _, md := range cd.middlewares {
			handler = md(handler)
		}

		return handler.HandleCommand(c)
	}

	panic(fmt.Sprintf("No handler registred for command '%s'", c.CommandType()))
}

func (cd *inProcessCommandDispatcher) UseMiddleware(mds ...CommandMiddleware) {
	cd.middlewares = append(cd.middlewares, mds...)
}

var onceCd sync.Once
var cd *inProcessCommandDispatcher

func InProcessCommandDispatcherInstance() *inProcessCommandDispatcher {
	onceCd.Do(func() {
		cd = &inProcessCommandDispatcher{make(map[string]CommandHandler), []CommandMiddleware{}}
	})

	return cd
}

