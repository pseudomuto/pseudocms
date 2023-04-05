package repo

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pseudomuto/pseudocms/pkg/models"
)

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

	var key interface{}
	if len(results) > 0 {
		key = fieldValue(*results[len(results)-1], o.orderBy)
	}

	return &ListResult[T]{
		Results:  results,
		LastKey:  key,
		LastPage: len(results) < o.pageSize,
	}, nil
}

func fieldValue[T Identifiable](obj T, field string) interface{} {
	v := reflect.ValueOf(obj)
	typ := v.Type()

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if f.Tag.Get("db") == field {
			return v.FieldByName(f.Name).Interface()
		}
	}

	// Special case for Model fields (embedded in most/all models).
	typ = reflect.TypeOf(models.Model{})
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if f.Tag.Get("db") == field {
			return v.FieldByName(f.Name).Interface()
		}
	}

	return nil
}
