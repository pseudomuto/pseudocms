package repo

import "github.com/gofrs/uuid"

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
