package cqrs

import (
	"github.com/google/uuid"
)

type Command interface {
	Id() uuid.UUID
	CorrelationID() uuid.UUID
	CausationID() uuid.UUID
	AggregateID() string
	CommandType() string
}
