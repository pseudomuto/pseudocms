package server

import (
	"github.com/gofrs/uuid"
	"github.com/pseudomuto/pseudocms/pkg/models"
)

// Repo describes an interface for working with database models. The purpose here
// is for testing. This is satisfied by models.Repo[T].
type Repo[T models.Identifiable] interface {
	Create(*T, models.CreateOptions) error
	Update(*T, models.CreateOptions) error
	Find(uuid.UUID, models.FindOptions) (*T, error)
}

// DefinitionsRepo is a typealias for a Definition repo.
type DefinitionsRepo Repo[models.Definition]

// FieldsRepo is a typealias for a Field repo.
type FieldsRepo Repo[models.Field]
