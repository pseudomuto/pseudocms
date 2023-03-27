package ctl

import (
	"strings"

	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/pseudomuto/pseudocms/pkg/ext"
	"github.com/spf13/cobra"
)

func fieldsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "fields",
		Short:   "Manage definition fields",
		Aliases: []string{"f"},
	}

	cmd.AddCommand(createFieldCmd())
	return cmd
}

func createFieldCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create <DEFINITION_ID>",
		Example: `
# Create a basic field for definition <def_id>.
pseudoctl fields create <def_id> \
  -n my_field \
  -d "some description" \
  -t text \
  -c "required, minLength(10)"`,
		Short:     "Create a new field",
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{"DEFINITION_ID"},
		RunE: func(cmd *cobra.Command, args []string) error {
			constraints, err := cmd.Flags().GetStringSlice("constraints")
			if err != nil {
				return err
			}

			// Trim the space around each of the values.
			constraints = ext.MapSlice(constraints, strings.TrimSpace)

			client := getAdminClient(cmd.Context())
			resp, err := client.CreateField(cmd.Context(), &v1.CreateFieldRequest{
				DefinitionId: args[0],
				Name:         cmd.Flags().Lookup("name").Value.String(),
				Description:  cmd.Flags().Lookup("description").Value.String(),
				FieldType:    toFieldType(cmd.Flags().Lookup("type").Value.String()),
				Constraints:  constraints,
			})

			if err != nil {
				return err
			}

			return printJSON(cmd, resp.Field)
		},
	}

	cmd.Flags().StringP("name", "n", "", "the name of the definition")
	cmd.Flags().StringP("description", "d", "", "the description of the definition")
	cmd.Flags().StringP("type", "t", "string", "the type of field")
	cmd.Flags().StringSliceP("constraints", "c", []string{}, "constraints for the field")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("description")

	return cmd
}

func toFieldType(typ string) v1.FieldType {
	switch strings.ToLower(typ) {
	case "float":
		return v1.FieldType_FIELD_TYPE_FLOAT
	case "int":
		return v1.FieldType_FIELD_TYPE_INT
	case "string":
		return v1.FieldType_FIELD_TYPE_STRING
	case "text":
		return v1.FieldType_FIELD_TYPE_TEXT
	}

	return v1.FieldType_FIELD_TYPE_UNSPECIFIED
}
