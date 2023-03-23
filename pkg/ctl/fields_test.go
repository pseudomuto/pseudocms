package ctl_test

import (
	"strings"

	"github.com/pseudomuto/pseudocms/pkg/models"
	"github.com/pseudomuto/pseudocms/pkg/testutil/factory"
	"github.com/pseudomuto/pseudocms/pkg/validation"
)

func (s *CtlSuite) TestCreateField() {
	def := factory.Definition.MustCreate().(models.Definition)
	s.Require().NoError(s.Conn().Create(&def))

	s.runCmd(
		"fields",
		"create",
		def.ID.String(),
		"-n", "description",
		"-d", "a description of the test field",
		"-t", string(models.Text),
		"-c", strings.Join([]string{validation.IsRequired().String(), validation.MinLength(3).String()}, ","),
	)
}
