package server

import (
	"net/http"

	"github.com/go-logr/logr"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// WithHTTPLogger is a middleware which emits log lines for each request.
func WithHTTPLogger(log logr.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)

		log.Info(
			"handled request",
			"path", r.URL.Path,
			"method", r.Method,
		)
	})
}

// GRPCLoggingInterceptors adds logging interceptors to all gRPC calls (unary and streaming).
func GRPCLoggingInterceptors(log *zap.Logger) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(log),
		),
		grpc.ChainStreamInterceptor(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(log),
		),
	}
}
