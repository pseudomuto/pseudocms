package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-logr/zapr"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"google.golang.org/grpc"
)

// ListenAndServe starts the gRPC server that serves API requests.
func ListenAndServe(addr string, opts ...Option) (chan<- os.Signal, <-chan bool) {
	svrOpts := makeOptions(opts)

	svr := grpc.NewServer(GRPCLoggingInterceptors(svrOpts.log)...)
	v1.RegisterHealthServiceServer(svr, HealthService())
	v1.RegisterAdminServiceServer(svr, AdminService(svrOpts.repoFactory))

	log := zapr.NewLogger(svrOpts.log)

	go func(svr *grpc.Server, opts *options) {
		signal.Notify(opts.sigTrap, svrOpts.sigs...)
		sig := <-opts.sigTrap
		log.Info("Received shutdown signal", "signal", sig)

		_, cancel := context.WithTimeout(context.Background(), opts.sdTimeout)
		defer cancel()

		svr.GracefulStop()
		close(opts.done)
	}(svr, svrOpts)

	go func(svr *grpc.Server, opts *options) {
		conn, err := net.Listen("tcp", addr)
		if err != nil {
			log.Error(err, "Failed to create TCP listener")
			close(svrOpts.done)
		}

		if err := svr.Serve(conn); err != nil {
			log.Error(err, "Failed to start/stop server cleanly")
		}
	}(svr, svrOpts)

	return svrOpts.sigTrap, svrOpts.done
}

// ListenAndServeHTTP starts the HTTP server that serves the gateway proxy.
func ListenAndServeHTTP(addr string, opts ...Option) (chan<- os.Signal, <-chan bool) {
	svrOpts := makeOptions(opts)
	log := zapr.NewLogger(svrOpts.log)

	handlers := []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1.RegisterAdminServiceHandlerFromEndpoint,
		v1.RegisterHealthServiceHandlerFromEndpoint,
	}

	ctx := context.Background()
	mux := runtime.NewServeMux()
	for _, h := range handlers {
		if err := h(ctx, mux, svrOpts.rpcHost, svrOpts.rpcDialOptions); err != nil {
			log.Error(err, "Failed to register service handler")
			close(svrOpts.done)
			return svrOpts.sigTrap, svrOpts.done
		}
	}

	svr := &http.Server{
		Addr:    addr,
		Handler: WithHTTPLogger(log, mux),
	}

	go func(svr *http.Server, opts *options) {
		signal.Notify(opts.sigTrap, svrOpts.sigs...)
		sig := <-opts.sigTrap
		log.Info("Received shutdown signal", "signal", sig)

		_, cancel := context.WithTimeout(context.Background(), opts.sdTimeout)
		defer cancel()

		// turn off keep alives for new connections
		svr.SetKeepAlivesEnabled(false)
		if err := svr.Shutdown(ctx); err != nil {
			log.Error(err, "Failed to shutdown gateway cleanly")
		}

		close(opts.done)
	}(svr, svrOpts)

	go func(svr *http.Server, opts *options) {
		if err := svr.ListenAndServe(); err != nil {
			log.Error(err, "Failed to start gateway server cleanly")
		}
	}(svr, svrOpts)

	return svrOpts.sigTrap, svrOpts.done
}
