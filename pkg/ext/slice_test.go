package ext_test

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/pseudomuto/pseudocms/pkg/ext"
	"github.com/stretchr/testify/require"
)

func TestMapSlice(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	exp := []string{"1", "2", "3", "4", "5"}
	require.Equal(t, exp, MapSlice(in, func(i int) string { return fmt.Sprintf("%d", i) }))

	in = []int{}
	exp = []string{}
	require.Equal(t, exp, MapSlice(in, func(i int) string { return fmt.Sprintf("%d", i) }))

	in = nil
	exp = []string{}
	require.Equal(t, exp, MapSlice(in, func(i int) string { return fmt.Sprintf("%d", i) }))

	t.Run("MapErr", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		res, err := MapSliceErr(in, func(i int) (string, error) { return fmt.Sprintf("%d", i), nil })
		require.NoError(t, err)
		require.Equal(t, []string{"1", "2", "3", "4", "5"}, res)

		res, err = MapSliceErr(in, func(i int) (string, error) { return "", errors.New("Boom") })
		require.Nil(t, res)
		require.EqualError(t, err, "Boom")
	})
}
