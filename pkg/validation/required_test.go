package validation_test

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/pseudomuto/pseudocms/pkg/validation"
	"github.com/stretchr/testify/require"
)

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
	require.Equal(t, fmt.Sprintf(`"%s"`, c.String()), string(res))
}
