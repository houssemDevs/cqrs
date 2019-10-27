package cqrs

type EventHandler interface {
	HandleEvent(ev Event)
}
