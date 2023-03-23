package ctl

import (
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/spf13/cobra"
)

func definitionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "definitions",
		Short:   "Manage content definitions",
		Aliases: []string{"d", "defs"},
	}

	cmd.AddCommand(createDefCmd())
	return cmd
}

func createDefCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create",
		Example: `# Create a new content defintion.
pseudoctl defs create \
  -n definition_name \
  -d "some description" \
`,
		Short: "Create a new content definition",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client := v1.NewAdminServiceClient(getClient())
			resp, err := client.CreateDefinition(cmd.Context(), &v1.CreateDefinitionRequest{
				Name:        cmd.Flags().Lookup("name").Value.String(),
				Description: cmd.Flags().Lookup("description").Value.String(),
			})
			if err != nil {
				return err
			}

			return printJSON(cmd, resp.Definition)
		},
	}

	cmd.Flags().StringP("name", "n", "", "the name of the definition")
	cmd.Flags().StringP("description", "d", "", "the description of the definition")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("description")

	return cmd
}
