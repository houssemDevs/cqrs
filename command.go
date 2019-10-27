package cqrs

import (
	"github.com/google/uuid"
)

type Command interface {
	CorrelationID() uuid.UUID
	CommandType() string
}


