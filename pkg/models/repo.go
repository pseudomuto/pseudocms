package models

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

// NewRepo creates a new Repo for T objects.
func NewRepo[T Identifiable](db *pop.Connection) *Repo[T] {
	return &Repo[T]{
		db: db,
	}
}

// Repo is a generic object repository for handling CRUD operations for
// Identifiable types.
type Repo[T Identifiable] struct {
	db *pop.Connection
}

// CreateOptions define some options for use when creating new entities.
type CreateOptions struct {
	// Eager, when true, ensures that all associations are created as well in the
	// same transaction.
	Eager bool
}

// Save creates the model in the database.
func (r *Repo[T]) Create(model *T, opts CreateOptions) error {
	return r.db.Transaction(func(tx *pop.Connection) error {
		if opts.Eager {
			tx = tx.Eager()
		}

		return tx.Create(model)
	})
}

// FindOptions define some options for executing Find queries.
type FindOptions struct {
	// Eager, when true, will load all associations for the record.
	Eager bool
}

// Find finds a record by ID.
func (r *Repo[T]) Find(id uuid.UUID, opts FindOptions) (*T, error) {
	db := r.db
	if opts.Eager {
		db = db.Eager()
	}

	var model T
	if err := db.Find(&model, id); err != nil {
		return nil, err
	}

	return &model, nil
}
