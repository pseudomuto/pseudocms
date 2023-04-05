package repo_test

import (
	"testing"

	"github.com/gobuffalo/pop/v6"
	"github.com/pseudomuto/pseudocms/pkg/testutil/testdb"
	tsuite "github.com/stretchr/testify/suite"
)

type suite struct {
	tsuite.Suite

	tdb  *testdb.TestDB
	conn *pop.Connection
}

func TestSuite(t *testing.T) {
	tsuite.Run(t, new(suite))
}

// SetupSuite ensures that the test database is created and a connection is
// available via Conn.
func (s *suite) SetupSuite() {
	s.tdb = testdb.New(s.T())

	conn, err := s.tdb.Open()
	s.Require().NoError(err)
	s.conn = conn
}

// TearDownSuite ensures the database connection is closed and that all temp
// resources are cleaned up.
func (s *suite) TearDownSuite() {
	s.conn.Close()
	s.tdb.Close()
}

// TearDownTest ensures that the entire database is truncated at the end of each
// test.
func (s *suite) TearDownTest() {
	s.Require().NoError(s.conn.TruncateAll())
}
