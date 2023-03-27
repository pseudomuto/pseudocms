package ctl_test

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
)

func (s *CtlSuite) TestHealthPing() {
	s.health.EXPECT().Ping(gomock.Any(), &v1.PingRequest{}).Return(
		&v1.PingResponse{Msg: "PONG"},
		nil,
	)

	s.runCmd("health", "ping")
}
