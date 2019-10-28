package cqrs

import (
	"fmt"
	"log"
	"sync"
)

type EventPublisher interface {
	RegisterListener(eh EventHandler)
	Publish(ev Event)
}

type eventPublisher struct {
	listeners map[string][]EventHandler
}

func (ep *eventPublisher) RegisterListener(eh EventHandler, events ...Event) {
	for _, event := range events {
		if handlers, ok := ep.listeners[event.EventType()]; ok {
			handlers = append(handlers, eh)
			ep.listeners[event.EventType()] = handlers
		}
	}
}

func (ep *eventPublisher) Publish(ev Event) {
	if handlers, ok := ep.listeners[ev.EventType()]; ok {
		for _, h := range handlers {
			h.HandleEvent(ev)
		}
	} else {
		log.Println(fmt.Sprintf("no handler for event '%s'", ev.EventType()))
	}
}

var onceEp sync.Once
var ep *eventPublisher

func InProcessEventPublisherInstance() *eventPublisher {
	onceEp.Do(func() {
		ep = &eventPublisher{make(map[string][]EventHandler)}
	})

	return ep
}



