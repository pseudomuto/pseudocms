package server_test

import (
	"context"
	"syscall"
	"testing"

	"github.com/go-logr/logr"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	. "github.com/pseudomuto/pseudocms/pkg/server"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestListenAndServe(t *testing.T) {
	sigs, done := ListenAndServe(
		"localhost:8091",
		WithLogger(logr.Discard()),
	)

	t.Cleanup(func() {
		sigs <- syscall.SIGTERM
		<-done
	})

	conn, err := grpc.Dial("localhost:8091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() { conn.Close() })

	client := v1.NewHealthServiceClient(conn)
	resp, err := client.Ping(context.Background(), new(v1.PingRequest))
	require.NoError(t, err)
	require.Equal(t, "PONG", resp.Msg)
}
