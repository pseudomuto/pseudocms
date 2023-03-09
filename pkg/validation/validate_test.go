package validation_test

import (
	"errors"
	"testing"

	. "github.com/pseudomuto/pseudocms/pkg/validation"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		given       interface{}
		constraints []Constraint
		want        []Result
	}{
		{
			given:       "a",
			constraints: []Constraint{IsRequired(), MinLength(2)},
			want: []Result{
				{Constraint: IsRequired(), Valid: true},
				{Constraint: MinLength(2), Valid: false},
			},
		},
		{
			given:       "abc",
			constraints: []Constraint{IsRequired(), MinLength(2), new(errorConstraint)},
			want: []Result{
				{Constraint: IsRequired(), Valid: true},
				{Constraint: MinLength(2), Valid: true},
				{Constraint: new(errorConstraint), Valid: false, Error: "boom"},
			},
		},
	}

	for _, tt := range tests {
		res := Validate(tt.given, tt.constraints...)

		require.Len(t, res, len(tt.want))
		for i, r := range tt.want {
			require.Equal(t, r, res[i])
		}
	}
}

type errorConstraint struct{}

func (e *errorConstraint) IsValid(interface{}) (bool, error) {
	return false, errors.New("boom")
}

func (e *errorConstraint) String() string { return "error" }
