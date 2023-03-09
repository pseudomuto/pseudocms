package testutil

import (
	"testing"

	"github.com/gobuffalo/pop/v6"
	"github.com/stretchr/testify/require"
)

// WithDB opens and closes a database connection to the test database.
func WithDB(t *testing.T, fn func(*pop.Connection)) {
	db, err := pop.Connect("test")
	require.NoError(t, err)

	t.Cleanup(func() {
		db.TruncateAll()
		db.Close()
	})

	fn(db)
}
