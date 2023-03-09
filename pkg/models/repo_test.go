package models_test

import (
	"testing"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	. "github.com/pseudomuto/pseudocms/pkg/models"
	"github.com/pseudomuto/pseudocms/pkg/testutil"
	"github.com/pseudomuto/pseudocms/pkg/testutil/factory"
	"github.com/stretchr/testify/require"
)

func TestRepoFind(t *testing.T) {
	testutil.WithDB(t, func(tx *pop.Connection) {
		def := factory.Definition.MustCreate().(Definition)
		require.NoError(t, tx.Create(&def))

		repo := NewRepo[Definition](tx)

		// when found
		d, err := repo.Find(def.ID, FindOptions{})
		require.NoError(t, err)
		require.Equal(t, def.Name, d.Name)

		// not found
		d, err = repo.Find(uuid.Must(uuid.NewV4()), FindOptions{})
		require.Nil(t, d)
		require.EqualError(t, err, "sql: no rows in result set")
	})
}

func TestRepoSave(t *testing.T) {
	t.Run("only root object", func(t *testing.T) {
		testutil.WithDB(t, func(tx *pop.Connection) {
			def := factory.Definition.MustCreate().(Definition)
			require.NotEmpty(t, def.Fields)

			repo := NewRepo[Definition](tx)
			repo.Create(&def, CreateOptions{})

			d, err := repo.Find(def.ID, FindOptions{Eager: true})
			require.NoError(t, err)
			require.Equal(t, def.Name, d.Name)
			require.Empty(t, d.Fields)
		})
	})

	t.Run("eager save", func(t *testing.T) {
		testutil.WithDB(t, func(tx *pop.Connection) {
			def := factory.Definition.MustCreate().(Definition)
			require.NotEmpty(t, def.Fields)

			repo := NewRepo[Definition](tx)
			require.NoError(t, repo.Create(&def, CreateOptions{Eager: true}))

			d, err := repo.Find(def.ID, FindOptions{Eager: true})
			require.NoError(t, err)
			require.Equal(t, def.Name, d.Name)
			require.Len(t, d.Fields, len(def.Fields))
		})
	})
}
