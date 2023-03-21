package server_test

import (
	"context"
	"testing"

	"github.com/gobuffalo/pop/v6"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/pseudomuto/pseudocms/pkg/models"
	. "github.com/pseudomuto/pseudocms/pkg/server"
	"github.com/pseudomuto/pseudocms/pkg/testutil"
	"github.com/stretchr/testify/require"
)

func TestAdminCreateDefinition(t *testing.T) {
	ctx := context.Background()

	testutil.WithDB(t, func(tx *pop.Connection) {
		repo := models.NewRepo[models.Definition](tx)
		svc := AdminService(repo)

		resp, err := svc.CreateDefinition(ctx, &v1.CreateDefinitionRequest{
			Name:        "Test Definition",
			Description: "Some Test Definition",
			Fields: []*v1.Field{
				{
					Name:        "Some Field",
					Description: "Some Field Description",
					FieldType:   v1.FieldType_FIELD_TYPE_STRING,
					Constraints: []string{"required"},
				},
			},
		})

		require.NoError(t, err)
		require.NotEmpty(t, resp.Definition.Id)
		require.NotEmpty(t, resp.Definition.Fields[0].Id)
	})
}
