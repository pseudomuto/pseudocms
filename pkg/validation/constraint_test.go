package validation_test

import (
	"testing"

	. "github.com/pseudomuto/pseudocms/pkg/validation"
	"github.com/stretchr/testify/require"
)

func TestParseConstraint(t *testing.T) {
	tests := []struct {
		given string
		want  Constraint
		err   string
	}{
		{given: "required", want: IsRequired()},
		{given: "maxLength(3)", want: MaxLength(3)},
		{given: "maxLength(ten)", err: "strconv.Atoi: parsing \"ten\": invalid syntax"},
		{given: "minLength(1)", want: MinLength(1)},
		{given: "minLength(ten)", err: "strconv.Atoi: parsing \"ten\": invalid syntax"},
		{given: "whodis", err: "unknown constraint: whodis"},
	}

	for _, tt := range tests {
		c, err := ParseConstraint(tt.given)
		require.Equal(t, tt.want, c, tt.given)
		if tt.err != "" {
			require.EqualError(t, err, tt.err, tt.given)
		}
	}
}
