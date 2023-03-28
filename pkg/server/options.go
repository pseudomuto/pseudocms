package server

import (
	"os"
	"syscall"
	"time"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type options struct {
	done           chan bool
	log            logr.Logger
	repoFactory    RepoFactory
	rpcDialOptions []grpc.DialOption
	rpcHost        string
	sigs           []os.Signal
	sigTrap        chan os.Signal
	sdTimeout      time.Duration
}

// Option describes a setup option for the HTTP server.
type Option interface {
	apply(*options)
}

type optFunc func(*options)

func (f optFunc) apply(o *options) { f(o) }

// WithDialOptions sets the DialOptions used for connecting to the upstream RPC server.
func WithDialOptions(opts ...grpc.DialOption) Option {
	return optFunc(func(o *options) { o.rpcDialOptions = opts })
}

// WithLogger species the logger to use for logging. By default, this is a development zap logger.
func WithLogger(log logr.Logger) Option {
	return optFunc(func(o *options) { o.log = log })
}

// WithRepoFactory sets the RepoFactory used to communicate with the database.
func WithRepoFactory(rf RepoFactory) Option {
	return optFunc(func(o *options) { o.repoFactory = rf })
}

// WithRPCHost sets the upstream RPC server for the gateway proxy.
func WithRPCHost(host string) Option {
	return optFunc(func(o *options) { o.rpcHost = host })
}

// WithSignals specifies the signal that will trigger a shutdown. By default these are syscall.SIGINT and syscall.SIGTERM.
func WithSignals(sigs ...os.Signal) Option {
	return optFunc(func(o *options) { o.sigs = sigs })
}

// WithShutdownTimeout specified the duration to wait for a graceful shutdown. Default is 5s.
func WithShutdownTimeout(d time.Duration) Option {
	return optFunc(func(o *options) { o.sdTimeout = d })
}

func makeOptions(svrOptions []Option) *options {
	opts := &options{
		done:      make(chan bool),
		log:       defaultLogger(),
		sigs:      []os.Signal{syscall.SIGINT, syscall.SIGTERM},
		sigTrap:   make(chan os.Signal),
		sdTimeout: 5 * time.Second,
	}

	for _, opt := range svrOptions {
		opt.apply(opts)
	}

	return opts
}

func defaultLogger() logr.Logger {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic("unable to initialize logger:" + err.Error())
	}

	return zapr.NewLogger(log)
}
