package factory

import (
	"time"

	"github.com/bluele/factory-go/factory"
	"github.com/gofrs/uuid"
	"github.com/pseudomuto/pseudocms/pkg/models"
)

// Model is a factory for generating base models.
var Model = factory.NewFactory(models.Model{}).
	Attr("ID", func(factory.Args) (interface{}, error) {
		return uuid.Must(uuid.NewV4()), nil
	}).
	Attr("CreatedAt", func(factory.Args) (interface{}, error) {
		return time.Now().UTC(), nil
	}).
	Attr("UpdatedAt", func(factory.Args) (interface{}, error) {
		return time.Now().UTC(), nil
	})
