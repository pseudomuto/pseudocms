package repo_test

import (
	"github.com/pseudomuto/pseudocms/pkg/ext"
	"github.com/pseudomuto/pseudocms/pkg/models"
	. "github.com/pseudomuto/pseudocms/pkg/repo"
	"github.com/pseudomuto/pseudocms/pkg/testutil/factory"
)

func (s *suite) TestUpdate() {
	kinds := func(fields []models.Field) []models.FieldKind {
		return ext.MapSlice(fields, func(f models.Field) models.FieldKind { return f.Kind })
	}

	def := factory.Definition.MustCreate().(models.Definition)
	s.Require().NotEmpty(def.Fields)

	repo := New[models.Definition](s.conn)
	s.Require().NoError(repo.Create(&def, CreateOptions{Eager: true}))

	def.Name = "Updated name"
	def.Fields[0].Kind = models.Integer
	s.Require().NoError(repo.Update(&def, UpdateOptions{}))

	d, err := repo.Find(def.ID, FindOptions{Eager: true})
	s.Require().NoError(err)
	s.Require().Equal(def.Name, d.Name)
	s.Require().NotContains(kinds(d.Fields), models.Integer)
}
