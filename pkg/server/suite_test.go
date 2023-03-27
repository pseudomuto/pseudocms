package server_test

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/pseudomuto/pseudocms/pkg/server"
	tsuite "github.com/stretchr/testify/suite"
)

type suite struct {
	tsuite.Suite

	ctrl  *gomock.Controller
	repos *mockRepoFactory
}

func TestSuite(t *testing.T) {
	tsuite.Run(t, new(suite))
}

func (s *suite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repos = &mockRepoFactory{
		defs:   NewMockDefinitionsRepo(s.ctrl),
		fields: NewMockFieldsRepo(s.ctrl),
	}
}

func (s *suite) TearDownTest() {
	s.ctrl.Finish()
}

type mockRepoFactory struct {
	defs   *MockDefinitionsRepo
	fields *MockFieldsRepo
}

func (m *mockRepoFactory) Definitions() server.DefinitionsRepo {
	return m.defs
}

func (m *mockRepoFactory) Fields() server.FieldsRepo {
	return m.fields
}
