package ctl_test

//go:generate go run github.com/golang/mock/mockgen -destination=mocks_test.go -package=ctl_test github.com/pseudomuto/pseudocms/pkg/api/v1 AdminServiceClient,HealthServiceClient

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"regexp"
	"testing"

	"github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
	. "github.com/pseudomuto/pseudocms/pkg/ctl"
	"github.com/pseudomuto/pseudocms/pkg/testutil"
	"github.com/stretchr/testify/suite"
)

var (
	//go:embed testdata/*.yaml
	fs embed.FS

	// uuidRegexp matches (roughly) a UUID
	uuidRegexp = regexp.MustCompile(`[\w]{8}-[\w]{4}-[\w]{4}-[\w]{4}-[\w]{12}`)
)

type CtlSuite struct {
	suite.Suite

	ctrl   *gomock.Controller
	admin  *MockAdminServiceClient
	health *MockHealthServiceClient
}

func TestCtl(t *testing.T) {
	suite.Run(t, new(CtlSuite))
}

func (s *CtlSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.admin = NewMockAdminServiceClient(s.ctrl)
	s.health = NewMockHealthServiceClient(s.ctrl)
}

func (s *CtlSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *CtlSuite) runCmd(in io.Reader, cmd ...string) {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	s.Require().NoError(Run(cmd, Options{
		AdminClient:  s.admin,
		HealthClient: s.health,
		In:           in,
		Out:          stdout,
		Err:          stderr,
	}))

	out := uuidRegexp.ReplaceAllString(stdout.String(), uuid.Nil.String())
	err := uuidRegexp.ReplaceAllString(stderr.String(), uuid.Nil.String())

	testutil.AssertGolden(
		s.T(),
		fmt.Sprintf("STDOUT:\n%s\n\nSTDERR:\n%s", out, err),
	)
}
