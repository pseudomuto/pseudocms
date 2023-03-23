package ctl_test

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/gofrs/uuid"
	. "github.com/pseudomuto/pseudocms/pkg/ctl"
	"github.com/pseudomuto/pseudocms/pkg/testutil"
)

// uuidRegexp matches (roughly) a UUID
var uuidRegexp = regexp.MustCompile(`[\w]{8}-[\w]{4}-[\w]{4}-[\w]{4}-[\w]{12}`)

func (s *CtlSuite) TestHelp() {
	s.runCmd("-h")
}

func (s *CtlSuite) TestVersion() {
	s.runCmd("-v")
}

func (s *CtlSuite) runCmd(cmd ...string) {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	cmd = append([]string{"--server", fmt.Sprintf("localhost:%d", s.port)}, cmd...)
	s.Require().NoError(Run(cmd, Options{Out: stdout, Err: stderr}))

	out := uuidRegexp.ReplaceAllString(stdout.String(), uuid.Nil.String())
	err := uuidRegexp.ReplaceAllString(stderr.String(), uuid.Nil.String())

	testutil.AssertGolden(
		s.T(),
		fmt.Sprintf("STDOUT:\n%s\n\nSTDERR:\n%s", out, err),
	)
}
