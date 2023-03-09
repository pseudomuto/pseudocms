package models

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
)

// Definition describes the details and metadata about content. It can be thought
// of like a schema for content.
type Definition struct {
	Model       `json:",inline"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Fields      []Field `json:"fields" has_many:"fields", order_by:"name ASC"`
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
