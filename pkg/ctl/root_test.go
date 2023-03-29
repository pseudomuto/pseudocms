package ctl_test

func (s *CtlSuite) TestHelp() {
	s.runCmd(nil, "-h")
}

func (s *CtlSuite) TestVersion() {
	s.runCmd(nil, "-v")
}
