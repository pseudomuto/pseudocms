package content_test

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/pseudomuto/pseudocms/pkg/content"
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

func TestIsRequired(t *testing.T) {
	tests := []struct {
		name  string
		given interface{}
		want  bool
	}{
		{name: "nil"},
		{name: "empty string", given: ""},
		{name: "white space", given: "  ", want: true},
		{name: "zero", given: 0},
		{name: "zero (uint)", given: uint(0)},
		{name: "zero (int32)", given: int32(0)},
		{name: "zero (uint32)", given: uint32(0)},
		{name: "zero (int64)", given: int64(0)},
		{name: "zero (uint64)", given: uint64(0)},
		{name: "zero", given: 0},
		{name: "zero (float32)", given: float32(0.0)},
		{name: "zero (float)", given: 0.0},
		{name: "custom type", given: IsRequired(), want: true},
	}

	c := IsRequired()
	for _, tt := range tests {
		res, err := c.IsValid(tt.given)
		require.Equal(t, tt.want, res, tt.name)
		require.NoError(t, err, tt.name)
	}

	res, err := json.Marshal(c)
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf(`"%s"`, c.Name()), string(res))
}

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
