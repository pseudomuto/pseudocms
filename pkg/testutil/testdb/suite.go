package testdb

import "github.com/stretchr/testify/suite"

// Suite is a test suite that creates a new ephemeral cockroach db at the beginning
// of the suite and shuts it down at the end.
type Suite struct {
	suite.Suite
	db *TestDB
}

func (s *Suite) SetupSuite() {
}

func (s *Suite) TearDownSuite() {
	s.db.Close()
}
