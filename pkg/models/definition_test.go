package models_test

import (
	"strings"
	"testing"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/pseudomuto/pseudocms/pkg/ext"
	. "github.com/pseudomuto/pseudocms/pkg/models"
	"github.com/pseudomuto/pseudocms/pkg/testutil"
	"github.com/pseudomuto/pseudocms/pkg/testutil/factory"
	"github.com/stretchr/testify/require"
)

func TestDefinitionFromProto(t *testing.T) {
	proto := &v1.Definition{
		Id:          uuid.Must(uuid.NewV4()).String(),
		Name:        "test",
		Description: "Some test definition",
		Fields: []*v1.Field{
			{
				Id:          uuid.Must(uuid.NewV4()).String(),
				Name:        "test",
				Description: "Some test field",
				FieldType:   v1.FieldType_FIELD_TYPE_STRING,
				Constraints: []string{"required", "minLength(3)"},
			},
		},
	}

	fields := ext.MapSlice(proto.Fields, func(f *v1.Field) Field { return *FieldFromProto(f) })

	def := DefinitionFromProto(proto)
	require.Equal(t, proto.Id, def.ID.String())
	require.Equal(t, proto.Name, def.Name)
	require.Equal(t, proto.Description, def.Description)
	require.Equal(t, fields, def.Fields)
}

func TestDefinitionToProto(t *testing.T) {
	def := factory.Definition.MustCreate().(Definition)
	fields := ext.MapSlice(def.Fields, func(f Field) *v1.Field { return f.ToProto() })

	proto := def.ToProto()
	require.Equal(t, def.ID.String(), proto.Id)
	require.Equal(t, def.Name, proto.Name)
	require.Equal(t, def.Description, proto.Description)
	require.Equal(t, fields, proto.Fields)
}

func TestDefinitionValidate(t *testing.T) {
	mkDef := func(kvs ...interface{}) Definition {
		opts := make(map[string]interface{})
		for i := 0; i < len(kvs)-1; i++ {
			opts[kvs[i].(string)] = kvs[i+1]
		}

		return factory.Definition.MustCreateWithOption(opts).(Definition)
	}

	tests := []struct {
		def    Definition
		errors map[string]string
	}{
		{
			def: mkDef("Name", ""),
			errors: map[string]string{
				"name": "Name can not be blank.",
			},
		},
		{
			def: mkDef("Name", strings.Repeat("a", 101)),
			errors: map[string]string{
				"name": "Name not in range(0, 100)",
			},
		},
		{
			def: mkDef("Description", ""),
			errors: map[string]string{
				"description": "Description can not be blank.",
			},
		},
	}

	testutil.WithDB(t, func(db *pop.Connection) {
		def := factory.Definition.MustCreate().(Definition)
		errs, err := def.Validate(db)
		require.Zero(t, errs.Count())
		require.NoError(t, err)

		for _, tt := range tests {
			errs, err := tt.def.Validate(db)
			require.NoError(t, err)
			require.Equal(t, len(tt.errors), errs.Count())

			for field, message := range tt.errors {
				require.Equal(t, []string{message}, errs.Get(field))
			}
		}
	})
}
