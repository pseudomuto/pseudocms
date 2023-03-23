package ctl_test

import (
	"fmt"
	"math/rand"
	"os"
	"syscall"
	"testing"

	"github.com/go-logr/logr"
	"github.com/pseudomuto/pseudocms/pkg/server"
	"github.com/pseudomuto/pseudocms/pkg/testutil/testdb"
	"github.com/stretchr/testify/suite"
)

type CtlSuite struct {
	testdb.Suite

	sigs chan<- os.Signal
	done <-chan bool
	port int
}

func TestCtl(t *testing.T) {
	suite.Run(t, new(CtlSuite))
}

func (s *CtlSuite) SetupSuite() {
	s.Suite.SetupSuite()

	// random port in the range [8000, 9000)
	s.port = rand.Intn(9000-8000+1) + 8000
	s.sigs, s.done = server.ListenAndServe(
		fmt.Sprintf("localhost:%d", s.port),
		server.WithDatabase(s.Conn()),
		server.WithLogger(logr.Discard()),
	)
}

func (s *CtlSuite) TearDownSuite() {
	s.sigs <- syscall.SIGTERM
	<-s.done

	s.Suite.TearDownSuite()
}
