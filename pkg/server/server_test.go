package server_test

import (
	"context"
	"io"
	"net/http"
	"syscall"
	"testing"

	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	. "github.com/pseudomuto/pseudocms/pkg/server"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestListenAndServe(t *testing.T) {
	withRPCServer(t, func(host string) {
		conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
		require.NoError(t, err)
		t.Cleanup(func() { conn.Close() })

		client := v1.NewHealthServiceClient(conn)
		resp, err := client.Ping(context.Background(), new(v1.PingRequest))
		require.NoError(t, err)
		require.Equal(t, "PONG", resp.Msg)
	})
}

func TestListenAndServeHTTP(t *testing.T) {
	withRPCServer(t, func(host string) {
		sigs, done := ListenAndServeHTTP(
			"localhost:8192",
			WithLogger(zaptest.NewLogger(t)),
			WithRPCHost(host),
			WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
		)

		t.Cleanup(func() {
			sigs <- syscall.SIGTERM
			<-done
		})

		resp, err := http.Get("http://localhost:8192/v1/health/ping")
		require.NoError(t, err)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		t.Cleanup(func() { resp.Body.Close() })

		require.Equal(t, `{"msg":"PONG"}`, string(body))
	})
}

func withRPCServer(t *testing.T, fn func(string)) {
	t.Helper()

	sigs, done := ListenAndServe(
		"localhost:8191",
		WithLogger(zaptest.NewLogger(t)),
	)

	t.Cleanup(func() {
		sigs <- syscall.SIGTERM
		<-done
	})

	fn("localhost:8191")
}
