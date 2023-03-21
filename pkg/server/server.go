package server

import (
	"context"
	"net"
	"os"
	"os/signal"

	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/pseudomuto/pseudocms/pkg/models"
	"google.golang.org/grpc"
)

// ListenAndServe starts the gRPC server that serves API requests.
func ListenAndServe(addr string, opts ...Option) (chan<- os.Signal, <-chan bool) {
	svrOpts := makeOptions(opts)

	svr := grpc.NewServer()
	v1.RegisterAdminServiceServer(svr, AdminService(models.NewRepo[models.Definition](svrOpts.db)))
	v1.RegisterHealthServiceServer(svr, HealthService())

	go func(svr *grpc.Server, opts *options) {
		signal.Notify(opts.sigTrap, svrOpts.sigs...)
		sig := <-opts.sigTrap
		opts.log.Info("Received shutdown signal", "signal", sig)

		_, cancel := context.WithTimeout(context.Background(), opts.sdTimeout)
		defer cancel()

		svr.GracefulStop()
		close(opts.done)
	}(svr, svrOpts)

	go func(svr *grpc.Server, opts *options) {
		conn, err := net.Listen("tcp", addr)
		if err != nil {
			opts.log.Error(err, "Failed to create TCP listener")
			close(svrOpts.done)
		}

		if err := svr.Serve(conn); err != nil {
			opts.log.Error(err, "Failed to start/stop server cleanly")
		}
	}(svr, svrOpts)

	return svrOpts.sigTrap, svrOpts.done
}
