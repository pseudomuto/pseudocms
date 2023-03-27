package server

//go:generate go run github.com/golang/mock/mockgen -destination=mocks_test.go -package=server_test . DefinitionsRepo,FieldsRepo

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

// RepoFactory describes an interface for creating repos as needed.
type RepoFactory interface {
	Definitions() DefinitionsRepo
	Fields() FieldsRepo
}
