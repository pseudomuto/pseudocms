package testutil

import (
	"testing"

	"github.com/gobuffalo/pop/v6"
	"github.com/pseudomuto/pseudocms/pkg/testutil/testdb"
	"github.com/stretchr/testify/require"
)

// WithDB opens and closes a database connection to the test database.
func WithDB(t *testing.T, fn func(*pop.Connection)) {
	db := testdb.New(t)
	conn, err := db.Open()
	require.NoError(t, err)

	t.Cleanup(db.Close)
	fn(conn)
}
