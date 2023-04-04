package models_test

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/pseudomuto/pseudocms/pkg/ext"
	. "github.com/pseudomuto/pseudocms/pkg/models"
	"github.com/pseudomuto/pseudocms/pkg/testutil/factory"
	"github.com/stretchr/testify/require"
)

func (s *suite) TestRepoCreate() {
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

	s.T().Run("eager", func(t *testing.T) {
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

func (s *suite) TestRepoUpdate() {
	kinds := func(fields []Field) []FieldKind {
		return ext.MapSlice(fields, func(f Field) FieldKind { return f.Kind })
	}

	def := factory.Definition.MustCreate().(Definition)
	s.Require().NotEmpty(def.Fields)

	repo := NewRepo[Definition](s.conn)
	s.Require().NoError(repo.Create(&def, CreateOptions{Eager: true}))

	def.Name = "Updated name"
	def.Fields[0].Kind = Integer
	s.Require().NoError(repo.Update(&def, UpdateOptions{}))

	d, err := repo.Find(def.ID, FindOptions{Eager: true})
	s.Require().NoError(err)
	s.Require().Equal(def.Name, d.Name)
	s.Require().NotContains(kinds(d.Fields), Integer)
}

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

func (s *suite) TestRepoList() {
	repo := NewRepo[Definition](s.conn)
	for i := 0; i < 10; i++ {
		def := factory.Definition.MustCreate().(Definition)
		s.Require().NoError(repo.Create(&def, CreateOptions{Eager: true}))
	}

	res, err := repo.List(Eager(true), PageSize(4), OrderBy("id", "asc"))
	s.Require().NoError(err)
	s.Require().Len(res.Results, 4)
	s.Require().False(res.LastPage)

	res, err = repo.List(Eager(true), PageSize(4), OrderBy("id", "asc"), AfterKey(res.LastKey))
	s.Require().NoError(err)
	s.Require().Len(res.Results, 4)
	s.Require().False(res.LastPage)

	res, err = repo.List(Eager(true), PageSize(4), OrderBy("id", "asc"), AfterKey(res.LastKey))
	s.Require().NoError(err)
	s.Require().Len(res.Results, 2)
	s.Require().True(res.LastPage)
}
