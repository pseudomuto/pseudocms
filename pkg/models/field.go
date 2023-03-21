package models

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/pop/v6/slices"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/pseudomuto/pseudocms/pkg/validation"
)

// FieldKind defines the different kinds of fields.
type FieldKind string

// Available field types
const (
	Float   FieldKind = "float"
	Integer FieldKind = "integer"
	String  FieldKind = "string"
	Text    FieldKind = "text"
)

var (
	fieldKindProtoMap = map[FieldKind]v1.FieldType{
		Float:   v1.FieldType_FIELD_TYPE_FLOAT,
		Integer: v1.FieldType_FIELD_TYPE_INT,
		String:  v1.FieldType_FIELD_TYPE_STRING,
		Text:    v1.FieldType_FIELD_TYPE_TEXT,
	}
)

// FieldKindFromProto creates a FieldKind from the supplied v1.FieldType.
func FieldKindFromProto(ft v1.FieldType) FieldKind {
	for k, v := range fieldKindProtoMap {
		if v == ft {
			return k
		}
	}

	panic("Invalid v1.FieldType supplied")
}

// ToProto converts this FieldKind to v1.FieldType.
func (fk FieldKind) ToProto() v1.FieldType {
	if v, ok := fieldKindProtoMap[fk]; ok {
		return v
	}

	return v1.FieldType_FIELD_TYPE_UNSPECIFIED
}

// Field defines a field.
type Field struct {
	Model        `json:",inline"`
	Definition   *Definition `belongs_to:"definition"`
	DefinitionID uuid.UUID   `json:"definitionId" db:"definition_id"`
	Name         string      `json:"name" db:"name"`
	Description  string      `json:"description" db:"description"`
	Kind         FieldKind   `json:"kind" db:"kind"`

	// Constraints can be a bit finicky. Best to avoid interaction with
	// this directly and use AddConstraint and GetConstraints.
	Constraints slices.String `json:"constraints" db:"constraints"`
}

// FieldFromProto creates a Field object from the supplied v1.Field.
func FieldFromProto(f *v1.Field) *Field {
	field := &Field{
		Name:        f.Name,
		Description: f.Description,
		Kind:        FieldKindFromProto(f.FieldType),
		Constraints: slices.String(f.Constraints),
	}

	if f.Id != "" {
		field.ID = uuid.Must(uuid.FromString(f.Id))
	}

	return field
}

// AddConstraint adds a new constraint to this field.
func (f *Field) AddConstraint(constraints ...validation.Constraint) {
	for _, c := range constraints {
		f.Constraints = append(f.Constraints, c.String())
	}
}

// GetConstraints returns the parsed constraints for this field.
func (f *Field) GetConstraints() ([]validation.Constraint, error) {
	constraints := make([]validation.Constraint, len(f.Constraints))
	for i, cs := range f.Constraints {
		c, err := validation.ParseConstraint(cs)
		if err != nil {
			return nil, err
		}

		constraints[i] = c
	}

	return constraints, nil
}

// IsValid determines whether or not the supplied value is valid according to the
// field's constraints.
func (f *Field) IsValid(v interface{}) ([]validation.Result, error) {
	constraints, err := f.GetConstraints()
	if err != nil {
		return nil, err
	}

	return validation.Validate(v, constraints...), nil
}

// ToProto converts this field to a v1.Field object.
func (f *Field) ToProto() *v1.Field {
	return &v1.Field{
		Id:          f.ID.String(),
		Name:        f.Name,
		Description: f.Description,
		FieldType:   f.Kind.ToProto(),
		Constraints: f.Constraints,
	}
}

// Validate is used to validate the Field object. This method is called
// automatically by pop when performing validation during saves.
func (f *Field) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: f.Name, Name: "Name"},
		&validators.StringLengthInRange{Field: f.Name, Name: "Name", Max: 100},
		&validators.StringIsPresent{Field: f.Description, Name: "Description"},
		&validators.StringInclusion{Field: string(f.Kind), Name: "Kind", List: []string{
			string(Float), string(Integer), string(String), string(Text),
		}},
	), nil
}
