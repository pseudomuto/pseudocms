package repo_test

import (
	"github.com/gofrs/uuid"
	"github.com/pseudomuto/pseudocms/pkg/models"
	. "github.com/pseudomuto/pseudocms/pkg/repo"
	"github.com/pseudomuto/pseudocms/pkg/testutil/factory"
)

func (s *suite) TestFind() {
	def := factory.Definition.MustCreate().(models.Definition)
	s.Require().NoError(s.conn.Create(&def))

	repo := New[models.Definition](s.conn)

	// when found
	d, err := repo.Find(def.ID, FindOptions{})
	s.Require().NoError(err)
	s.Require().Equal(def.Name, d.Name)

	// not found
	d, err = repo.Find(uuid.Must(uuid.NewV4()), FindOptions{})
	s.Require().Nil(d)
	s.Require().EqualError(err, "sql: no rows in result set")
}
