package models_test

import (
	"testing"

	"github.com/pseudomuto/pseudocms/pkg/testutil/testdb"
	tsuite "github.com/stretchr/testify/suite"
)

type suite struct {
	testdb.Suite
}

func TestSuite(t *testing.T) {
	tsuite.Run(t, new(suite))
}
