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
