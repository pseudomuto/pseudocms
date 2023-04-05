package repo

import "github.com/gobuffalo/pop/v6"

// CreateOptions define some options for use when creating new entities.
type CreateOptions struct {
	// Eager, when true, ensures that all associations are created as well in the
	// same transaction.
	Eager bool
}

// Create creates the model in the database.
func (r *Repo[T]) Create(model *T, opts CreateOptions) error {
	return r.db.Transaction(func(tx *pop.Connection) error {
		if opts.Eager {
			tx = tx.Eager()
		}

		return tx.Create(model)
	})
}
