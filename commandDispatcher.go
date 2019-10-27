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

type commandDispatcher struct {
	handlers map[string]CommandHandler
	middlewares []CommandMiddleware
}

func (cd *commandDispatcher) RegisterHandler(handler CommandHandler, c Command, mds ...CommandMiddleware) {
	for _, md := range mds {
		handler = md(handler)
	}

	if registeredHandler, found := cd.handlers[c.CommandType()]; found {
		panic(fmt.Sprintf("A handler '%s' is already registred for command '%s'", reflect.TypeOf(registeredHandler), c.CommandType()))
	}

	cd.handlers[c.CommandType()] = handler
}

func (cd *commandDispatcher) Dispatch(c Command) error {
	if handler, found := cd.handlers[c.CommandType()]; found {
		for _, md := range cd.middlewares {
			handler = md(handler)
		}

		return handler.HandleCommand(c)
	}

	panic(fmt.Sprintf("No handler registred for command '%s'", c.CommandType()))
}

func (cd *commandDispatcher) UseMiddleware(mds ...CommandMiddleware) {
	cd.middlewares = append(cd.middlewares, mds...)
}

var onceCd sync.Once
var cd *commandDispatcher

func CommandDispatcherInstance() *commandDispatcher {
	onceCd.Do(func() {
		cd = &commandDispatcher{make(map[string]CommandHandler), []CommandMiddleware{}}
	})

	return cd
}

