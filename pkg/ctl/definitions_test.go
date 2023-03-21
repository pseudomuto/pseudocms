package ctl_test

func (s *CtlSuite) TestCreateDefinition() {
	s.runCmd(
		"definitions",
		"create",
		"-n", "test definition",
		"-d", "a description of the test definition",
	)
}
