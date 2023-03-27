package models_test

import (
	"testing"

	"github.com/gofrs/uuid"
	. "github.com/pseudomuto/pseudocms/pkg/models"
	"github.com/pseudomuto/pseudocms/pkg/testutil/factory"
	"github.com/stretchr/testify/require"
)

func (s *suite) TestRepoFind() {
	def := factory.Definition.MustCreate().(Definition)
	s.Require().NoError(s.conn.Create(&def))

	repo := NewRepo[Definition](s.conn)

	// when found
	d, err := repo.Find(def.ID, FindOptions{})
	s.Require().NoError(err)
	s.Require().Equal(def.Name, d.Name)

	// not found
	d, err = repo.Find(uuid.Must(uuid.NewV4()), FindOptions{})
	s.Require().Nil(d)
	s.Require().EqualError(err, "sql: no rows in result set")
}

func (s *suite) TestRepoSave() {
	s.T().Run("only root object", func(t *testing.T) {
		def := factory.Definition.MustCreate().(Definition)
		require.NotEmpty(t, def.Fields)

		repo := NewRepo[Definition](s.conn)
		repo.Create(&def, CreateOptions{})

		d, err := repo.Find(def.ID, FindOptions{Eager: true})
		require.NoError(t, err)
		require.Equal(t, def.Name, d.Name)
		require.Empty(t, d.Fields)
	})

	s.T().Run("eager save", func(t *testing.T) {
		def := factory.Definition.MustCreate().(Definition)
		require.NotEmpty(t, def.Fields)

		repo := NewRepo[Definition](s.conn)
		require.NoError(t, repo.Create(&def, CreateOptions{Eager: true}))

		d, err := repo.Find(def.ID, FindOptions{Eager: true})
		require.NoError(t, err)
		require.Equal(t, def.Name, d.Name)
		require.Len(t, d.Fields, len(def.Fields))
	})
}
