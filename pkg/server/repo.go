package server

import "github.com/pseudomuto/pseudocms/pkg/models"

// Repo describes an interface for working with database models. The purpose here
// is for testing. This is satisfied by models.Repo[T].
type Repo[T models.Identifiable] interface {
	Create(*T, models.CreateOptions) error
}

// DefinitionsRepo is a typealias for a Definition repo.
type DefinitionsRepo Repo[models.Definition]
