package server_test

//go:generate go run github.com/golang/mock/mockgen -destination=mocks_api_test.go -package=server_test github.com/pseudomuto/pseudocms/pkg/api/v1 AdminService_ListDefinitionsServer

import (
	"context"
	"sort"
	"testing"

	"github.com/gobuffalo/pop/v6/slices"
	"github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/pseudomuto/pseudocms/pkg/models"
	"github.com/pseudomuto/pseudocms/pkg/repo"
	. "github.com/pseudomuto/pseudocms/pkg/server"
	"github.com/pseudomuto/pseudocms/pkg/testutil/factory"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *suite) TestAdminCreateDefinition() {
	ctx := context.Background()
	svc := AdminService(s.repos)

	s.repos.defs.EXPECT().Create(&models.Definition{
		Name:        "Test Definition",
		Description: "Some Test Definition",
		Fields: []models.Field{
			{
				Name:        "Some Field",
				Description: "Some Field Description",
				Kind:        models.String,
				Constraints: slices.String{"required"},
			},
		},
	}, repo.CreateOptions{Eager: true}).Return(nil)

	resp, err := svc.CreateDefinition(ctx, &v1.CreateDefinitionRequest{
		Name:        "Test Definition",
		Description: "Some Test Definition",
		Fields: []*v1.Field{
			{
				Name:        "Some Field",
				Description: "Some Field Description",
				FieldType:   v1.FieldType_FIELD_TYPE_STRING,
				Constraints: []string{"required"},
			},
		},
	})

	s.Require().NoError(err)
	s.Require().NotEmpty(resp.Definition.Id)
	s.Require().NotEmpty(resp.Definition.Fields[0].Id)
}

func (s *suite) TestAdminGetDefinition() {
	ctx := context.Background()
	svc := AdminService(s.repos)

	id := uuid.Must(uuid.NewV4())
	s.repos.defs.EXPECT().Find(id, repo.FindOptions{Eager: true}).Return(
		&models.Definition{
			Model:       models.Model{ID: id},
			Name:        "Some Definition",
			Description: "Some Test Description",
			Fields: []models.Field{
				{
					Model:       models.Model{ID: uuid.Must(uuid.NewV4())},
					Name:        "Some Field",
					Description: "Some Field Description",
					Kind:        models.Text,
					Constraints: []string{"required"},
				},
			},
		},
		nil,
	)

	resp, err := svc.GetDefinition(ctx, &v1.GetDefinitionRequest{Id: id.String()})
	s.Require().NoError(err)
	s.Require().NotNil(resp.Definition)
}

func (s *suite) TestAdminListDefinitions() {
	svc := AdminService(s.repos)

	defs := make([]*models.Definition, 5)
	stream := NewMockAdminService_ListDefinitionsServer(s.ctrl)

	for i := 0; i < len(defs); i++ {
		d, ok := factory.Definition.MustCreate().(models.Definition)
		s.Require().True(ok)
		defs[i] = &d
	}

	// Stable sort by ID.
	sort.Slice(defs, func(i, j int) bool { return defs[i].ID.String() < defs[j].ID.String() })

	for _, d := range defs {
		// One expectation per definition.
		stream.EXPECT().Send(d.ToProto()).Return(nil)
	}

	s.repos.defs.EXPECT().List(gomock.Any()).Return(&repo.ListResult[models.Definition]{
		LastKey:  defs[len(defs)-1].ID,
		LastPage: true,
		Results:  defs,
	}, nil)

	err := svc.ListDefinitions(&v1.ListDefinitionsRequest{}, stream)
	s.Require().NoError(err)

	s.T().Run("single page", func(t *testing.T) {
		// Expect indices [1, 3]
		for i := 1; i < 4; i++ {
			stream.EXPECT().Send(defs[i].ToProto()).Return(nil)
		}

		s.repos.defs.EXPECT().List(gomock.Any()).Return(&repo.ListResult[models.Definition]{
			LastKey:  defs[3].ID,
			LastPage: false,
			Results:  defs[1:4], // range [i, N)
		}, nil)

		err := svc.ListDefinitions(&v1.ListDefinitionsRequest{
			AfterKey:   defs[0].ID.String(),
			MaxResults: wrapperspb.Int32(3),
		}, stream)
		s.Require().NoError(err)
	})
}

func (s *suite) TestAdminCreateField() {
	ctx := context.Background()
	svc := AdminService(s.repos)

	id := uuid.Must(uuid.NewV4())
	s.repos.fields.EXPECT().Create(&models.Field{
		DefinitionID: id,
		Name:         "Some Field",
		Description:  "Some Field Description",
		Kind:         models.String,
		Constraints:  slices.String{"required", "minLength(3)"},
	}, repo.CreateOptions{}).Return(nil)

	resp, err := svc.CreateField(ctx, &v1.CreateFieldRequest{
		DefinitionId: id.String(),
		Name:         "Some Field",
		Description:  "Some Field Description",
		FieldType:    v1.FieldType_FIELD_TYPE_STRING,
		Constraints:  []string{"required", "minLength(3)"},
	})
	s.Require().NoError(err)
	s.Require().NotEmpty(resp.Field.Id)
}
