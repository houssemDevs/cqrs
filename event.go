package cqrs

import "github.com/google/uuid"

type Event interface {
	CausationID() uuid.UUID
	EventType() string
}
