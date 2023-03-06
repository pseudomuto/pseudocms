package content_test

import (
	"encoding/json"
	"testing"

	. "github.com/pseudomuto/pseudocms/pkg/content"
	"github.com/stretchr/testify/require"
)

func TestFieldIsValid(t *testing.T) {
	field := Field{
		Name:        "MyField",
		Description: "Some details about the field",
		Kind:        String,
		Constraints: []Constraint{
			IsRequired(),
			MinLength(3),
			MaxLength(10),
		},
	}

	tests := []struct {
		name  string
		given interface{}
		want  []ValidationResult
	}{
		{
			name:  "valid value",
			given: "my name",
			want: []ValidationResult{
				{Constraint: "required", Valid: true},
				{Constraint: "minLength(3)", Valid: true},
				{Constraint: "maxLength(10)", Valid: true},
			},
		},
		{
			name:  "blank value",
			given: "",
			want: []ValidationResult{
				{Constraint: "required", Valid: false},
				{Constraint: "minLength(3)", Valid: false},
				{Constraint: "maxLength(10)", Valid: true},
			},
		},
		{
			name:  "too short",
			given: "my",
			want: []ValidationResult{
				{Constraint: "required", Valid: true},
				{Constraint: "minLength(3)", Valid: false},
				{Constraint: "maxLength(10)", Valid: true},
			},
		},
		{
			name:  "too long",
			given: "my name is far too long",
			want: []ValidationResult{
				{Constraint: "required", Valid: true},
				{Constraint: "minLength(3)", Valid: true},
				{Constraint: "maxLength(10)", Valid: false},
			},
		},
		{
			name:  "wrong type",
			given: 1,
			want: []ValidationResult{
				{Constraint: "required", Valid: true},
				{Constraint: "minLength(3)", Valid: false, Error: "invalid type for minLength"},
				{Constraint: "maxLength(10)", Valid: false, Error: "invalid type for maxLength"},
			},
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.want, field.IsValid(tt.given), tt.name)
	}
}

func TestFieldSerde(t *testing.T) {
	field := Field{
		Name:        "MyField",
		Description: "Some details about the field",
		Kind:        String,
		Constraints: []Constraint{
			IsRequired(),
			MinLength(3),
			MaxLength(10),
		},
	}

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

		require.EqualError(
			t,
			json.Unmarshal(data, &field),
			"unknown constraint: unknown",
		)
	})
}
