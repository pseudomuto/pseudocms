package server_test

import (
	"context"

	"github.com/gobuffalo/pop/v6/slices"
	"github.com/gofrs/uuid"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/pseudomuto/pseudocms/pkg/models"
	. "github.com/pseudomuto/pseudocms/pkg/server"
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
	}, models.CreateOptions{Eager: true}).Return(nil)

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
	}, models.CreateOptions{}).Return(nil)

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
