package models

import (
	"fmt"
	"strings"

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

// Create creates the model in the database.
func (r *Repo[T]) Create(model *T, opts CreateOptions) error {
	return r.db.Transaction(func(tx *pop.Connection) error {
		if opts.Eager {
			tx = tx.Eager()
		}

		return tx.Create(model)
	})
}

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

type listOpts struct {
	eager    bool
	afterKey interface{}
	orderBy  string
	order    string
	pageSize int
}

// ListOption defines an option for fetching multiple records.
type ListOption interface {
	apply(*listOpts)
}

type listOptFunc func(*listOpts)

func (f listOptFunc) apply(o *listOpts) { f(o) }

// AfterKey sets the last seen key for pagination. This is typically the value
// returned by a list call via the LastKey field.
func AfterKey(k interface{}) ListOption {
	return listOptFunc(func(o *listOpts) { o.afterKey = k })
}

// Eager determines whether to load assocications or not.
func Eager(eager bool) ListOption {
	return listOptFunc(func(o *listOpts) { o.eager = eager })
}

// OrderBy sets the sort order. This is important for proper pagination.
func OrderBy(f, dir string) ListOption {
	return listOptFunc(func(o *listOpts) {
		o.orderBy = f
		o.order = dir
	})
}

// PageSize sets the maximum number of rows to be returned by a List call.
func PageSize(n int) ListOption {
	return listOptFunc(func(o *listOpts) { o.pageSize = n })
}

// ListResult encapsulates a single page of results returned from a List call.
type ListResult[T Identifiable] struct {
	// The page of results
	Results []*T

	// The last key for this list.
	LastKey interface{}

	// LastPage indicates whether or not this represents the last page of results.
	// It isn't perfect and will be true when len(Results) < the page size.
	LastPage bool
}

// List returns a paged list of entities from the database using the supplied options.
//
// By default the first 100 rows, sorted by created_at (oldest to newest) will
// be returned.
func (r *Repo[T]) List(opts ...ListOption) (*ListResult[T], error) {
	o := &listOpts{
		orderBy:  "created_at",
		order:    "asc",
		pageSize: 100,
	}

	for _, opt := range opts {
		opt.apply(o)
	}

	db := r.db
	if o.eager {
		db = db.Eager()
	}

	qry := db.
		Order(fmt.Sprintf("%s %s", o.orderBy, o.order)).
		Limit(o.pageSize)

	if o.afterKey != nil {
		op := ">"
		if strings.ToLower(o.order) == "desc" {
			op = "<"
		}

		qry = qry.Where(fmt.Sprintf("%s %s ?", o.orderBy, op), o.afterKey)
	}

	var results []*T
	if err := qry.All(&results); err != nil {
		return nil, err
	}

	return &ListResult[T]{
		Results:  results,
		LastKey:  (*results[len(results)-1]).GetID(),
		LastPage: len(results) < o.pageSize,
	}, nil
}
