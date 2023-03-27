package main

import (
	"log"
	"os"

	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/pseudomuto/pseudocms/pkg/ctl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	host := os.Getenv("PSEUDOCMS_SERVER")
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	if err := ctl.Run(os.Args[1:], ctl.Options{
		AdminClient:  v1.NewAdminServiceClient(conn),
		HealthClient: v1.NewHealthServiceClient(conn),
	}); err != nil {
		log.Fatal(err)
	}
}
