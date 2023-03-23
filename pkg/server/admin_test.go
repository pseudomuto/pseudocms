package server_test

import (
	"context"

	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/pseudomuto/pseudocms/pkg/models"
	. "github.com/pseudomuto/pseudocms/pkg/server"
)

func (s *suite) TestAdminCreateDefinition() {
	ctx := context.Background()
	repo := models.NewRepo[models.Definition](s.Conn())
	svc := AdminService(repo, nil)

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
	svc := AdminService(
		models.NewRepo[models.Definition](s.Conn()),
		models.NewRepo[models.Field](s.Conn()),
	)

	defResp, err := svc.CreateDefinition(ctx, &v1.CreateDefinitionRequest{
		Name:        "Test Definition",
		Description: "Some Test Definition",
	})
	s.Require().NoError(err)
	s.Require().Empty(defResp.Definition.Fields)

	resp, err := svc.CreateField(ctx, &v1.CreateFieldRequest{
		DefinitionId: defResp.Definition.Id,
		Name:         "Some Field",
		Description:  "Some Field Description",
		FieldType:    v1.FieldType_FIELD_TYPE_STRING,
		Constraints:  []string{"required", "minLength(3)"},
	})
	s.Require().NoError(err)
	s.Require().NotEmpty(resp.Field.Id)
}
