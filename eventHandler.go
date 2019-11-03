package cqrs

type EventHandler interface {
	HandleEvent(ev Event)
}

type EventHandlerFunc func(ev Event)

func (ehf EventHandlerFunc) HandleEvent(ev Event) {
	ehf(ev)
}
