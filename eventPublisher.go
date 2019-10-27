package cqrs

import "sync"

type EventPublisher interface {
	RegisterListener(eh EventHandler)
	Publish(ev Event)
}

type eventPublisher struct {
	listeners []EventHandler
}

func (ep *eventPublisher) RegisterListener(eh EventHandler) {
	ep.listeners = append(ep.listeners, eh)
}

func (ep *eventPublisher) Publish(ev Event) {
	for _, eh := range ep.listeners {
		eh.HandleEvent(ev)
	}
}

var onceEp sync.Once
var ep *eventPublisher

func EventPublisherInstance() *eventPublisher {
	onceEp.Do(func() {
		ep = &eventPublisher{[]EventHandler{}}
	})

	return ep
}



