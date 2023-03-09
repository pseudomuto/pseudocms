package factory

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"github.com/pseudomuto/pseudocms/pkg/models"
)

// Definition is a factory for creating Definition objects.
var Definition = factory.NewFactory(models.Definition{}).
	Attr("Model", func(factory.Args) (interface{}, error) {
		return Model.MustCreate(), nil
	}).
	Attr("Name", func(factory.Args) (interface{}, error) {
		return randomdata.Noun(), nil
	}).
	Attr("Description", func(factory.Args) (interface{}, error) {
		return randomdata.Paragraph(), nil
	}).
	Attr("Fields", func(factory.Args) (interface{}, error) {
		return []models.Field{
			Field.MustCreate().(models.Field),
			Field.MustCreate().(models.Field),
			Field.MustCreate().(models.Field),
		}, nil
	})
