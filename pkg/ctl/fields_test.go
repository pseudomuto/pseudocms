package ctl_test

import (
	"github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
)

func (s *CtlSuite) TestCreateField() {
	defID := uuid.Must(uuid.FromString("f0b86eb0-4db0-4a31-8180-30ba65a7bcc8"))
	fieldID := uuid.Must(uuid.FromString("1139ac6b-cff8-47c5-8341-0345cefdb371"))

	req := &v1.CreateFieldRequest{
		DefinitionId: defID.String(),
		Name:         "description",
		Description:  "a description of the test field",
		FieldType:    v1.FieldType_FIELD_TYPE_STRING,
		Constraints:  []string{"required", "minLength(3)"},
	}

	s.admin.EXPECT().CreateField(gomock.Any(), req).Return(
		&v1.CreateFieldResponse{
			Field: &v1.Field{
				Id:          fieldID.String(),
				Name:        req.Name,
				Description: req.Description,
				FieldType:   req.FieldType,
				Constraints: req.Constraints,
			},
		},
		nil,
	)

	s.runCmd(
		"fields",
		"create",
		defID.String(),
		"-n", req.Name,
		"-d", req.Description,
		"-t", "string",
		"-c", "required, minLength(3)",
	)
}
