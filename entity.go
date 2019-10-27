package cqrs

type Entity interface {
	Id() string
	Changes() []Event
	Errors() []error
	CommitChanges()
	ClearErrors()
}
