package cqrs

import "github.com/google/uuid"

type Query interface {
	CorrelationID() uuid.UUID
	QueryType() string
}
