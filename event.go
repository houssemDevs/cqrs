package cqrs

import "github.com/google/uuid"

type Event interface {
	Id() uuid.UUID
	CorrelationID() uuid.UUID
	CausationID() uuid.UUID
	AggregateID() string
	EventType() string
}
