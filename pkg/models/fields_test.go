package models_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/gobuffalo/pop/v6/slices"
	. "github.com/pseudomuto/pseudocms/pkg/models"
	"github.com/pseudomuto/pseudocms/pkg/testutil/factory"
	"github.com/pseudomuto/pseudocms/pkg/validation"
	"github.com/stretchr/testify/require"
)

func TestFieldAddConstraint(t *testing.T) {
	field := Field{}
	require.Empty(t, field.Constraints)

	field.AddConstraint(validation.IsRequired(), validation.MinLength(3))
	require.Equal(t, field.Constraints, slices.String{"required", "minLength(3)"})
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
