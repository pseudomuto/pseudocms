package validation_test

import (
	"testing"

	. "github.com/pseudomuto/pseudocms/pkg/validation"
	"github.com/stretchr/testify/require"
)

func TestMinLength(t *testing.T) {
	tests := []struct {
		name  string
		given interface{}
		want  bool
	}{
		{name: "too short", given: "ab"},
		{name: "exact", given: "abc", want: true},
		{name: "long", given: "abcd", want: true},
	}

	c := MinLength(3)
	for _, tt := range tests {
		res, err := c.IsValid(tt.given)
		require.NoError(t, err, tt.name)
		require.Equal(t, tt.want, res, tt.name)
	}

	res, err := c.IsValid(0) // invalid type
	require.False(t, res)
	require.EqualError(t, err, "invalid type for minLength")
}

func TestMaxLength(t *testing.T) {
	tests := []struct {
		name  string
		given interface{}
		want  bool
	}{
		{name: "too short", given: "ab", want: true},
		{name: "exact", given: "abc", want: true},
		{name: "long", given: "abcd", want: false},
	}

	c := MaxLength(3)
	for _, tt := range tests {
		res, err := c.IsValid(tt.given)
		require.NoError(t, err, tt.name)
		require.Equal(t, tt.want, res, tt.name)
	}

	res, err := c.IsValid(0) // invalid type
	require.False(t, res)
	require.EqualError(t, err, "invalid type for maxLength")
}
