package ctl

import (
	"io"
	"sync"

	"github.com/pseudomuto/pseudocms"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client     grpc.ClientConnInterface
	clientSync sync.Once
)

type Options struct {
	In  io.Reader
	Err io.Writer
	Out io.Writer
}

// Run executes
func Run(args []string, opts Options) error {
	cmd := &cobra.Command{
		Use:     "pseudoctl",
		Short:   "A tool for working with pseudocms",
		Version: pseudocms.Version(),
	}

	if opts.Out != nil {
		cmd.SetOut(opts.Out)
	}

	// Will be available on all subcommands.
	cmd.PersistentFlags().StringP("server", "s", "", "The server to connect to [env: PSEUDOCMS_SERVER]")

	// Automatically bind PSEUDOCMS_* env vars
	viper.SetEnvPrefix("pseudocms")
	viper.AutomaticEnv()

	// Ensure flags override env vars
	viper.BindPFlag("server", cmd.PersistentFlags().Lookup("server"))

	cmd.AddCommand(definitionsCmd(), healthCmd())
	cmd.SetArgs(args)
	return cmd.Execute()
}

func getClient() grpc.ClientConnInterface {
	clientSync.Do(func() {
		host := viper.GetString("server")
		conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}

		client = conn
	})

	return client
}
