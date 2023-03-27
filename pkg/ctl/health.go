package ctl

import (
	"fmt"

	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/spf13/cobra"
)

func healthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "health",
		Short: "Query server health",
	}

	cmd.AddCommand(pingCmd())
	return cmd
}

func pingCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "Ping the server to make sure it's up",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := getHealthClient(cmd.Context())
			resp, err := client.Ping(cmd.Context(), new(v1.PingRequest))
			if err != nil {
				return err
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), resp.Msg)
			return err
		},
	}
}
