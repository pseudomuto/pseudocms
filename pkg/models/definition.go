package models

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/pseudomuto/pseudocms/pkg/ext"
)

// Definition describes the details and metadata about content. It can be thought
// of like a schema for content.
type Definition struct {
	Model       `json:",inline"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Fields      []Field `json:"fields" has_many:"fields", order_by:"name ASC"`
}

// DefinitionFromProto creates a new Definition based on the supplied v1.Definition.
func DefinitionFromProto(d *v1.Definition) *Definition {
	def := &Definition{
		Name:        d.Name,
		Description: d.Description,
		Fields:      ext.MapSlice(d.Fields, func(f *v1.Field) Field { return *FieldFromProto(f) }),
	}

	if d.Id != "" {
		def.ID = uuid.Must(uuid.FromString(d.Id))
	}

	return def
}

// ToProto converts this Definition into a v1.Definition.
func (d *Definition) ToProto() *v1.Definition {
	return &v1.Definition{
		Id:          d.ID.String(),
		Name:        d.Name,
		Description: d.Description,
		Fields:      ext.MapSlice(d.Fields, func(f Field) *v1.Field { return f.ToProto() }),
	}
}

// Validate validates the Definition object before saving it to the database.
// This function is called automatically by pop.
func (d *Definition) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: d.Name, Name: "Name"},
		&validators.StringLengthInRange{Field: d.Name, Name: "Name", Max: 100},
		&validators.StringIsPresent{Field: d.Description, Name: "Description"},
	), nil
}
