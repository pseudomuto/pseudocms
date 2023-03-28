package ctl_test

import (
	"github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
)

func (s *CtlSuite) TestCreateDefinition() {
	defID := uuid.Must(uuid.FromString("f0b86eb0-4db0-4a31-8180-30ba65a7bcc8"))

	req := &v1.CreateDefinitionRequest{
		Name:        "test defintion",
		Description: "description of test definition",
	}

	s.admin.EXPECT().CreateDefinition(gomock.Any(), req).Return(
		&v1.CreateDefinitionResponse{
			Definition: &v1.Definition{
				Id:          defID.String(),
				Name:        req.Name,
				Description: req.Description,
			},
		},
		nil,
	)

	s.runCmd(
		"definitions",
		"create",
		"-n", req.Name,
		"-d", req.Description,
	)
}

func (s *CtlSuite) TestGetDefinition() {
	defID := uuid.Must(uuid.FromString("f0b86eb0-4db0-4a31-8180-30ba65a7bcc8"))

	s.admin.EXPECT().
		GetDefinition(gomock.Any(), &v1.GetDefinitionRequest{Id: defID.String()}).
		Return(&v1.GetDefinitionResponse{
			Definition: &v1.Definition{
				Id:          defID.String(),
				Name:        "test",
				Description: "Some test definition",
				Fields: []*v1.Field{
					{
						Id:          uuid.Must(uuid.NewV4()).String(),
						Name:        "test",
						Description: "Some test field",
						FieldType:   v1.FieldType_FIELD_TYPE_STRING,
						Constraints: []string{"required", "minLength(3)"},
					},
				},
			},
		}, nil)

	s.runCmd(
		"definitions",
		"get",
		defID.String(),
	)
}
