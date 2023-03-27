package ctl_test

func (s *CtlSuite) TestHelp() {
	s.runCmd("-h")
}

func (s *CtlSuite) TestVersion() {
	s.runCmd("-v")
}
