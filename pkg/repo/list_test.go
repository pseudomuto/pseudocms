package repo_test

import (
	"math"
	"sort"

	"github.com/gofrs/uuid"
	"github.com/pseudomuto/pseudocms/pkg/ext"
	"github.com/pseudomuto/pseudocms/pkg/models"
	. "github.com/pseudomuto/pseudocms/pkg/repo"
	"github.com/pseudomuto/pseudocms/pkg/testutil/factory"
)

func (s *suite) TestList() {
	repo := New[models.Definition](s.conn)
	defs := make([]models.Definition, 10)
	for i := 0; i < len(defs); i++ {
		defs[i] = factory.Definition.MustCreate().(models.Definition)
		s.Require().NoError(repo.Create(&defs[i], CreateOptions{Eager: true}))
	}

	s.Run("full list", func() {
		res, err := repo.List(PageSize(len(defs) + 10))
		s.Require().NoError(err)
		s.Require().Len(res.Results, len(defs))
		s.Require().True(res.LastPage)
	})

	s.Run("simple pagination", func() {
		pageSize := int(math.Ceil(float64(len(defs)) / 3.0))
		opts := []ListOption{Eager(true), PageSize(pageSize), OrderBy("id", "asc")}

		// keep a list of sorted ids for verification of LastKey below
		ids := ext.MapSlice(defs, func(d models.Definition) uuid.UUID { return d.ID })
		sort.Slice(ids, func(i, j int) bool { return ids[i].String() < ids[j].String() })

		res, err := repo.List(opts...)
		s.Require().NoError(err)
		s.Require().Len(res.Results, pageSize)
		s.Require().Equal(ids[pageSize-1], res.LastKey)
		s.Require().False(res.LastPage) // false negative, but avoids querying count

		res, err = repo.List(append(opts, AfterKey(res.LastKey))...)
		s.Require().NoError(err)
		s.Require().Len(res.Results, pageSize)
		s.Require().Equal(ids[pageSize*2-1], res.LastKey)
		s.Require().False(res.LastPage)

		res, err = repo.List(append(opts, AfterKey(res.LastKey))...)
		s.Require().NoError(err)
		s.Require().Less(len(res.Results), pageSize)
		s.Require().Equal(ids[len(ids)-1], res.LastKey)
		s.Require().True(res.LastPage)

		// No more pages!
		res, err = repo.List(append(opts, AfterKey(res.LastKey))...)
		s.Require().NoError(err)
		s.Require().Nil(res.LastKey)
		s.Require().True(res.LastPage)
		s.Require().Zero(len(res.Results))
	})

	s.Run("sort descending", func() {
		// ids in descending order
		ids := ext.MapSlice(defs, func(d models.Definition) uuid.UUID { return d.ID })
		sort.Slice(ids, func(i, j int) bool { return ids[i].String() > ids[j].String() })

		pageSize := int(math.Ceil(float64(len(defs)) / 2.0))

		res, err := repo.List(PageSize(pageSize), OrderBy("id", "desc"))
		s.Require().NoError(err)
		s.Require().Len(res.Results, pageSize)
		s.Require().Equal(ids[pageSize-1], res.LastKey)
		s.Require().False(res.LastPage)

		for i, id := range ids {
			if i == pageSize {
				break
			}

			s.Require().Equal(id, res.Results[i].ID)
		}

		res, err = repo.List(PageSize(pageSize), OrderBy("id", "desc"), AfterKey(res.LastKey))
		s.Require().NoError(err)
		s.Require().Len(res.Results, len(ids)-pageSize)
		s.Require().Equal(ids[len(ids)-1], res.LastKey)
		s.Require().False(res.LastPage)

		for i, id := range ids {
			if i < pageSize {
				continue
			}

			s.Require().Equal(id, res.Results[i-pageSize].ID)
		}
	})

	s.Run("sort by custom field", func() {
		// expNames in order
		expNames := ext.MapSlice(defs, func(d models.Definition) string { return d.Name })
		sort.Slice(expNames, func(i, j int) bool { return expNames[i] < expNames[j] })

		res, err := repo.List(PageSize(len(defs)), OrderBy("name", "asc"))
		s.Require().NoError(err)
		s.Require().Len(res.Results, len(defs))
		s.Require().Equal(expNames[len(expNames)-1], res.LastKey)

		names := ext.MapSlice(res.Results, func(d *models.Definition) string { return d.Name })
		s.Require().Equal(expNames, names)

		// now descending
		sort.Slice(expNames, func(i, j int) bool { return expNames[i] > expNames[j] })
		res, err = repo.List(PageSize(len(defs)), OrderBy("name", "desc"))
		s.Require().NoError(err)
		s.Require().Len(res.Results, len(defs))
		s.Require().Equal(expNames[len(expNames)-1], res.LastKey)

		names = ext.MapSlice(res.Results, func(d *models.Definition) string { return d.Name })
		s.Require().Equal(expNames, names)
	})
}
