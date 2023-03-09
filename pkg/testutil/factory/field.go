package factory

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"github.com/pseudomuto/pseudocms/pkg/models"
)

// Field is a factory for creating Field objects.
var Field = factory.NewFactory(models.Field{}).
	Attr("Model", func(factory.Args) (interface{}, error) {
		return Model.MustCreate(), nil
	}).
	Attr("Name", func(factory.Args) (interface{}, error) {
		return randomdata.Noun(), nil
	}).
	Attr("Description", func(factory.Args) (interface{}, error) {
		return randomdata.Paragraph(), nil
	}).
	Attr("Kind", func(factory.Args) (interface{}, error) {
		return models.String, nil
	})
