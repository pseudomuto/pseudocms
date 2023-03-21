package ctl_test

func (s *CtlSuite) TestHealthPing() {
	s.runCmd("health", "ping")
}
