package testdb

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/stretchr/testify/suite"
)

// Suite is a Suite that constructs an ephemeral database to be used in all
// the tests on the suite.
//
// The database is created once at setup and then truncated after every test.
type Suite struct {
	suite.Suite

	tdb  *TestDB
	conn *pop.Connection
}

// SetupSuite ensures that the test database is created and a connection is
// available via Conn.
func (s *Suite) SetupSuite() {
	s.tdb = New(s.T())

	conn, err := s.tdb.Open()
	s.Require().NoError(err)
	s.conn = conn
}

// TearDownSuite ensures the database connection is closed and that all temp
// resources are cleaned up.
func (s *Suite) TearDownSuite() {
	s.conn.Close()
	s.tdb.Close()
}

// TearDownTest ensures that the entire database is truncated at the end of each
// test.
func (s *Suite) TearDownTest() {
	s.Require().NoError(s.conn.TruncateAll())
}

// Conn returns an open connection to the test database.
func (s *Suite) Conn() *pop.Connection {
	return s.conn
}
