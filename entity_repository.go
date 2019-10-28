package cqrs

type Repository interface {
	Load(id string) (Entity, error)
	Save(e Entity) error
}
