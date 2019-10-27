package cqrs

type CommandMiddleware func (handler CommandHandler) CommandHandler
