package repo

import "github.com/gobuffalo/pop/v6"

// UpdateOptions defines some options for use when updating existing entities.
type UpdateOptions CreateOptions

// Update updates the model in the database.
func (r *Repo[T]) Update(model *T, opts UpdateOptions) error {
	return r.db.Transaction(func(tx *pop.Connection) error {
		if opts.Eager {
			tx = tx.Eager()
		}

		return tx.Update(model)
	})
}
