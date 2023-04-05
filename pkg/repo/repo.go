package repo

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

// Identifiable describes an interface for types with an ID.
type Identifiable interface {
	// GetID returns the id for this thing.
	GetID() uuid.UUID
}

// New creates a new Repo for T objects.
func New[T Identifiable](db *pop.Connection) *Repo[T] {
	return &Repo[T]{
		db: db,
	}
}

// Repo is a generic object repository for handling CRUD operations for
// Identifiable types.
type Repo[T Identifiable] struct {
	db *pop.Connection
}
