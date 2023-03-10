package server

import (
	"context"

	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
)

// HealthService returns a new HealthServiceServer instance for handling health requests.
func HealthService() v1.HealthServiceServer {
	return new(healthService)
}

type healthService struct{}

// Ping simply returns PONG to signal that the server is up.
func (s *healthService) Ping(ctx context.Context, r *v1.PingRequest) (*v1.PingResponse, error) {
	return &v1.PingResponse{Msg: "PONG"}, nil
}
