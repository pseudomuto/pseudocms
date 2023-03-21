package models_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/gobuffalo/pop/v6/slices"
	"github.com/gofrs/uuid"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	. "github.com/pseudomuto/pseudocms/pkg/models"
	"github.com/pseudomuto/pseudocms/pkg/testutil/factory"
	"github.com/pseudomuto/pseudocms/pkg/validation"
	"github.com/stretchr/testify/require"
)

func TestFieldKindToProto(t *testing.T) {
	tests := map[FieldKind]v1.FieldType{
		Float:         v1.FieldType_FIELD_TYPE_FLOAT,
		Integer:       v1.FieldType_FIELD_TYPE_INT,
		String:        v1.FieldType_FIELD_TYPE_STRING,
		Text:          v1.FieldType_FIELD_TYPE_TEXT,
		FieldKind(""): v1.FieldType_FIELD_TYPE_UNSPECIFIED,
	}

	for k, v := range tests {
		require.Equal(t, v, k.ToProto())
	}
}

func TestFieldKindFromProto(t *testing.T) {
	tests := map[FieldKind]v1.FieldType{
		Float:   v1.FieldType_FIELD_TYPE_FLOAT,
		Integer: v1.FieldType_FIELD_TYPE_INT,
		String:  v1.FieldType_FIELD_TYPE_STRING,
		Text:    v1.FieldType_FIELD_TYPE_TEXT,
	}

	for k, v := range tests {
		require.Equal(t, k, FieldKindFromProto(v))
	}

	require.Panics(t, func() { _ = FieldKindFromProto(v1.FieldType_FIELD_TYPE_UNSPECIFIED) })
}

func TestFieldAddConstraint(t *testing.T) {
	field := Field{}
	require.Empty(t, field.Constraints)

	field.AddConstraint(validation.IsRequired(), validation.MinLength(3))
	require.Equal(t, field.Constraints, slices.String{"required", "minLength(3)"})
}

func TestFieldFromProto(t *testing.T) {
	proto := v1.Field{
		Id:          uuid.Must(uuid.NewV4()).String(),
		Name:        "test",
		Description: "Some test field",
		FieldType:   v1.FieldType_FIELD_TYPE_STRING,
		Constraints: []string{"required", "minLength(3)"},
	}

	field := FieldFromProto(&proto)
	require.Equal(t, proto.Id, field.ID.String())
	require.Equal(t, proto.Name, field.Name)
	require.Equal(t, proto.Description, field.Description)
	require.Equal(t, FieldKindFromProto(proto.FieldType), field.Kind)
	require.Equal(t, proto.Constraints, []string(field.Constraints))
}

func TestFieldToProto(t *testing.T) {
	field := factory.Field.MustCreate().(Field)
	field.AddConstraint(validation.IsRequired(), validation.MinLength(3))

	proto := field.ToProto()
	require.Equal(t, field.ID.String(), proto.Id)
	require.Equal(t, field.Name, proto.Name)
	require.Equal(t, field.Description, proto.Description)
	require.Equal(t, field.Kind.ToProto(), proto.FieldType)
	require.Equal(t, field.Constraints.Interface(), proto.Constraints)
}

func TestFieldValidate(t *testing.T) {
	mkField := func(kvs ...interface{}) Field {
		opts := make(map[string]interface{})
		for i := 0; i < len(kvs)-1; i++ {
			opts[kvs[i].(string)] = kvs[i+1]
		}

		return factory.Field.MustCreateWithOption(opts).(Field)
	}

	tests := []struct {
		field  Field
		errors map[string]string
	}{
		{
			field: mkField("Name", ""),
			errors: map[string]string{
				"name": "Name can not be blank.",
			},
		},
		{
			field: mkField("Name", strings.Repeat("a", 101)),
			errors: map[string]string{
				"name": "Name not in range(0, 100)",
			},
		},
		{
			field: mkField("Description", ""),
			errors: map[string]string{
				"description": "Description can not be blank.",
			},
		},
		{
			field: mkField("Kind", FieldKind("")),
			errors: map[string]string{
				"kind": "Kind is not in the list [float, integer, string, text].",
			},
		},
	}

	for _, tt := range tests {
		errs, err := tt.field.Validate(nil)
		require.NoError(t, err)
		require.Equal(t, len(tt.errors), errs.Count(), errs)

		for field, message := range tt.errors {
			require.Equal(t, []string{message}, errs.Get(field))
		}
	}
}

func TestFieldIsValid(t *testing.T) {
	field := factory.Field.MustCreateWithOption(map[string]interface{}{
		"Constraints": slices.String{"required", "minLength(3)", "maxLength(10)"},
	}).(Field)

	tests := []struct {
		name  string
		given interface{}
		want  []validation.Result
	}{
		{
			name:  "valid value",
			given: "my name",
			want: []validation.Result{
				{Constraint: validation.IsRequired(), Valid: true},
				{Constraint: validation.MinLength(3), Valid: true},
				{Constraint: validation.MaxLength(10), Valid: true},
			},
		},
		{
			name:  "blank value",
			given: "",
			want: []validation.Result{
				{Constraint: validation.IsRequired(), Valid: false},
				{Constraint: validation.MinLength(3), Valid: false},
				{Constraint: validation.MaxLength(10), Valid: true},
			},
		},
		{
			name:  "too short",
			given: "my",
			want: []validation.Result{
				{Constraint: validation.IsRequired(), Valid: true},
				{Constraint: validation.MinLength(3), Valid: false},
				{Constraint: validation.MaxLength(10), Valid: true},
			},
		},
		{
			name:  "too long",
			given: "my name is far too long",
			want: []validation.Result{
				{Constraint: validation.IsRequired(), Valid: true},
				{Constraint: validation.MinLength(3), Valid: true},
				{Constraint: validation.MaxLength(10), Valid: false},
			},
		},
		{
			name:  "wrong type",
			given: 1,
			want: []validation.Result{
				{Constraint: validation.IsRequired(), Valid: true},
				{Constraint: validation.MinLength(3), Valid: false, Error: "invalid type for minLength"},
				{Constraint: validation.MaxLength(10), Valid: false, Error: "invalid type for maxLength"},
			},
		},
	}

	for _, tt := range tests {
		res, err := field.IsValid(tt.given)
		require.NoError(t, err)
		require.Equal(t, tt.want, res, tt.name)
	}
}

func TestFieldSerde(t *testing.T) {
	field := factory.Field.MustCreateWithOption(map[string]interface{}{
		"Constraints": slices.String{"required", "minLength(3)", "maxLength(10)"},
	}).(Field)

	// serialize and deserialize to make sure we don't lose anything
	// in the process
	res, err := json.Marshal(field)
	require.NoError(t, err)

	var serdeField Field
	require.NoError(t, json.Unmarshal(res, &serdeField))
	require.Equal(t, field, serdeField)

	t.Run("invalid json", func(t *testing.T) {
		require.Contains(
			t,
			json.Unmarshal([]byte(`{"name": 0}`), &field).Error(),
			"json: cannot unmarshal number into Go struct",
		)
	})

	t.Run("unknown constraint", func(t *testing.T) {
		data := []byte(`{
			"name": "test",
			"constraints": ["required", "unknown"]
		}`)

		require.NoError(t, json.Unmarshal(data, &field))
		_, err := field.GetConstraints()
		require.EqualError(t, err, "unknown constraint: unknown")
	})
}
