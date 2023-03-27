package ctl

import (
	"encoding/json"
	"io"

	"github.com/pseudomuto/pseudocms"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Options defines options for managing I/O streams.
type Options struct {
	In  io.Reader
	Err io.Writer
	Out io.Writer

	AdminClient  v1.AdminServiceClient
	HealthClient v1.HealthServiceClient
}

// Run executes the pseudoctl cli command.
func Run(args []string, opts Options) error {
	cmd := &cobra.Command{
		Use:     "pseudoctl",
		Short:   "A tool for working with pseudocms",
		Version: pseudocms.Version(),
	}

	// set clients on the context
	cmd.SetContext(setHealthClient(
		setAdminClient(cmd.Context(), opts.AdminClient),
		opts.HealthClient,
	))

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

	cmd.AddCommand(definitionsCmd(), fieldsCmd(), healthCmd())
	cmd.SetArgs(args)
	return cmd.Execute()
}

// printJSON prints the marshaled obj using the stdout stream for the command.
func printJSON(cmd *cobra.Command, obj interface{}) error {
	res, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}

	_, err = cmd.OutOrStdout().Write(append(res, byte('\n')))
	return err
}
