package main

import (
	"log"
	"os"

	"github.com/pseudomuto/pseudocms/pkg/ctl"
)

func main() {
	if err := ctl.Run(os.Args[1:], ctl.Options{}); err != nil {
		log.Fatal(err)
	}
}
