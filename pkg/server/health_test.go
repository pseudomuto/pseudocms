package server_test

import (
	context "context"
	"testing"

	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	. "github.com/pseudomuto/pseudocms/pkg/server"
	"github.com/stretchr/testify/require"
)

func TestHealthPing(t *testing.T) {
	ctx := context.Background()
	resp, err := HealthService().Ping(ctx, new(v1.PingRequest))
	require.NoError(t, err)
	require.Equal(t, "PONG", resp.Msg)
}
