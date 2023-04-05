package repo_test

import (
	"testing"

	"github.com/pseudomuto/pseudocms/pkg/models"
	. "github.com/pseudomuto/pseudocms/pkg/repo"
	"github.com/pseudomuto/pseudocms/pkg/testutil/factory"
	"github.com/stretchr/testify/require"
)

func (s *suite) TestCreate() {
	s.T().Run("only root object", func(t *testing.T) {
		def := factory.Definition.MustCreate().(models.Definition)
		require.NotEmpty(t, def.Fields)

		repo := New[models.Definition](s.conn)
		repo.Create(&def, CreateOptions{})

		d, err := repo.Find(def.ID, FindOptions{Eager: true})
		require.NoError(t, err)
		require.Equal(t, def.Name, d.Name)
		require.Empty(t, d.Fields)
	})

	s.T().Run("eager", func(t *testing.T) {
		def := factory.Definition.MustCreate().(models.Definition)
		require.NotEmpty(t, def.Fields)

		repo := New[models.Definition](s.conn)
		require.NoError(t, repo.Create(&def, CreateOptions{Eager: true}))

		d, err := repo.Find(def.ID, FindOptions{Eager: true})
		require.NoError(t, err)
		require.Equal(t, def.Name, d.Name)
		require.Len(t, d.Fields, len(def.Fields))
	})
}
