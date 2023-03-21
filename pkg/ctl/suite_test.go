package ctl_test

import (
	"fmt"
	"math/rand"
	"os"
	"syscall"
	"testing"

	"github.com/go-logr/logr"
	"github.com/gobuffalo/pop/v6"
	"github.com/pseudomuto/pseudocms/pkg/server"
	"github.com/pseudomuto/pseudocms/pkg/testutil/testdb"
	"github.com/stretchr/testify/suite"
)

type CtlSuite struct {
	suite.Suite

	tdb  *testdb.TestDB
	conn *pop.Connection
	sigs chan<- os.Signal
	done <-chan bool
	port int
}

func TestCtl(t *testing.T) {
	suite.Run(t, new(CtlSuite))
}

func (s *CtlSuite) SetupSuite() {
	s.tdb = testdb.New(s.T())

	conn, err := s.tdb.Open()
	s.Require().NoError(err)
	s.conn = conn

	// random port in the range [8000, 9000)
	s.port = rand.Intn(9000-8000+1) + 8000
	s.sigs, s.done = server.ListenAndServe(
		fmt.Sprintf("localhost:%d", s.port),
		server.WithDatabase(s.conn),
		server.WithLogger(logr.Discard()),
	)
}

func (s *CtlSuite) TearDownSuite() {
	s.sigs <- syscall.SIGTERM
	<-s.done
	s.conn.Close()
	s.tdb.Close()
}
