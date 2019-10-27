package cqrs

type CommandHandler interface {
	HandleCommand(c Command) error
}

type CommandHandlerFunc func(c Command) error

func (chf CommandHandlerFunc) HandleCommand(c Command) error {
	return chf(c)
}
