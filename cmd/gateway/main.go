package main

import (
	"flag"
	"log"
	"strings"

	"github.com/pseudomuto/pseudocms/pkg/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := flag.String("addr", "localhost:9000", "the host address to bind to")
	rpcHost := flag.String("rpc-host", "", "the upstream RPC server to connect to")
	flag.Parse()

	zlog, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer zlog.Sync()

	*addr = strings.TrimSpace(*addr)
	*rpcHost = strings.TrimSpace(*rpcHost)
	if *rpcHost == "" {
		zlog.Fatal("no rpc-host specified")
	}

	_, done := server.ListenAndServeHTTP(
		*addr,
		server.WithLogger(zlog),
		server.WithRPCHost(*rpcHost),
		server.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	<-done
}
