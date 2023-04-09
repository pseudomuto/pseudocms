package ctl

import (
	"io"

	"github.com/gofrs/uuid"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/pseudomuto/pseudocms/pkg/ext"
	"github.com/pseudomuto/pseudocms/pkg/models"
	"github.com/spf13/cobra"
)

func definitionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "definitions",
		Short:   "Manage content definitions",
		Aliases: []string{"d", "defs"},
	}

	cmd.AddCommand(createDefCmd(), getDefCmd(), listDefsCmd(), updateDefCmd())
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
			def, err := parseStdinOrFile[models.Definition](cmd)
			if err != nil {
				return err
			}

			// Flags override values from JSON file (if supplied).
			def.Name = stringFlag(cmd, "name", def.Name)
			def.Description = stringFlag(cmd, "description", def.Description)

			client := getAdminClient(cmd.Context())
			resp, err := client.CreateDefinition(cmd.Context(), &v1.CreateDefinitionRequest{
				Name:        def.Name,
				Description: def.Description,
				Fields:      ext.MapSlice(def.Fields, func(f models.Field) *v1.Field { return f.ToProto() }),
			})
			if err != nil {
				return err
			}

			return printJSON(cmd, resp.Definition)
		},
	}

	cmd.Flags().StringP("file", "f", "", "a YAML file containing the definition")
	cmd.Flags().StringP("name", "n", "", "the name of the definition")
	cmd.Flags().StringP("description", "d", "", "the description of the definition")

	return cmd
}

func getDefCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <DEFINITION_ID>",
		Short: "Look up a definition by id",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := getAdminClient(cmd.Context())
			resp, err := client.GetDefinition(cmd.Context(), &v1.GetDefinitionRequest{
				Id: args[0],
			})

			if err != nil {
				return err
			}

			return printJSON(cmd, resp.Definition)
		},
	}
}

func listDefsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all definitions",
		RunE: func(cmd *cobra.Command, args []string) error {
			eager, err := cmd.Flags().GetBool("associations")
			if err != nil {
				return err
			}

			client := getAdminClient(cmd.Context())
			stream, err := client.ListDefinitions(cmd.Context(), &v1.ListDefinitionsRequest{
				Eager:   eager,
				OrderBy: "id",
			})

			if err != nil {
				return err
			}

			for {
				def, err := stream.Recv()
				if err != nil {
					if err == io.EOF {
						break
					}

					return err
				}

				if err := printJSON(cmd, def); err != nil {
					return err
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolP("associations", "a", false, "load associations (fields, etc).")
	return cmd
}

func updateDefCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "update <DEFINITION_ID>",
		Short:     "Update a definition",
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{"DEFINITION_ID"},
		RunE: func(cmd *cobra.Command, args []string) error {
			def, err := parseStdinOrFile[models.Definition](cmd)
			if err != nil {
				return err
			}

			// Flags override values from JSON file (if supplied).
			def.ID = uuid.Must(uuid.FromString(args[0]))
			def.Name = stringFlag(cmd, "name", def.Name)
			def.Description = stringFlag(cmd, "description", def.Description)

			client := getAdminClient(cmd.Context())
			resp, err := client.UpdateDefinition(cmd.Context(), &v1.UpdateDefinitionRequest{
				Id:          def.ID.String(),
				Name:        def.Name,
				Description: def.Description,
			})

			if err != nil {
				return err
			}

			return printJSON(cmd, resp.Definition)
		},
	}

	cmd.Flags().StringP("file", "f", "", "a YAML file containing the definition")
	cmd.Flags().StringP("name", "n", "", "the name of the definition")
	cmd.Flags().StringP("description", "d", "", "the description of the definition")

	return cmd
}
