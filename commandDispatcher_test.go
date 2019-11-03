package cqrs

import (
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

type BaseCommand struct {
	id            uuid.UUID
	correlationID uuid.UUID
	causationID   uuid.UUID
}

func NewBaseCommand() BaseCommand {
	uuid, _ := uuid.NewUUID()
	return BaseCommand{uuid, uuid, uuid}
}

func (c BaseCommand) CorrelationID() uuid.UUID {
	return c.correlationID
}

func (c BaseCommand) CausationID() uuid.UUID {
	return c.causationID
}

func (c BaseCommand) Id() uuid.UUID {
	return c.id
}

type AddTodo struct {
	BaseCommand
	aggregateID string
	Text        string
}

func NewAddTodo(text string) AddTodo {
	return AddTodo{NewBaseCommand(), "123", text}
}

func (c AddTodo) AggregateID() string {
	return c.aggregateID
}

func (c AddTodo) CommandType() string {
	return reflect.TypeOf(c).Name()
}

type DeleteTodo struct {
	BaseCommand
	aggregateId string
	Text        string
}

func NewDeleteTodo(text string) DeleteTodo {
	return DeleteTodo{NewBaseCommand(), "123", text}
}

func (c DeleteTodo) AggregateID() string {
	return c.aggregateId
}

func (c DeleteTodo) CommandType() string {
	return reflect.TypeOf(c).Name()
}

var AddTodoToList = CommandHandlerFunc(func(c Command) error {
	if c, ok := c.(AddTodo); ok {
		log.Println("Adding todo to list ", c.Text)
		return nil
	}
	return errors.New("Unknow command " + reflect.TypeOf(c).Name())
})

var LoggingMiddleware = func(h CommandHandler) CommandHandler {
	return CommandHandlerFunc(func(c Command) error {
		switch c.(type) {
		case AddTodo:
			log.Println("Middleware in action")
			return h.HandleCommand(c)
		case DeleteTodo:
			log.Println("Going to panic no handler")
			return h.HandleCommand(c)
		default:
			return nil
		}
	})
}

func TestInProcessCommandDispatcher(t *testing.T) {
	InProcessCommandDispatcherInstance().UseMiddleware(LoggingMiddleware)
	InProcessCommandDispatcherInstance().RegisterHandler(AddTodoToList, AddTodo{})

	t.Run("Correclty handling command", func(t *testing.T) {
		err := InProcessCommandDispatcherInstance().Dispatch(NewAddTodo("doing basic cqrs"))
		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("Panic if no handler", func(t *testing.T) {
		defer func() {
			if err := recover(); err == nil {
				t.Error("no handler and yet no panic")
			}
		}()
		InProcessCommandDispatcherInstance().Dispatch(NewDeleteTodo("doing basic cqrs"))
	})
}
