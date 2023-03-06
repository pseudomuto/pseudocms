package content

import "encoding/json"

// FieldKind defines the different kinds of fields.
type FieldKind string

// Available field types
const (
	Float   FieldKind = "float"
	Integer FieldKind = "integer"
	String  FieldKind = "string"
	Text    FieldKind = "text"
)

// field is effectively a clone of Field only where Constraints is a string slice.
// This is necessary since the serialized form of a constraint needs to be parsed
// in order to be the right type.
//
// See UnmarshalJSON for how this is handled.
type field struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Kind        FieldKind `json:"kind"`
	Constraints []string  `json:"constraints"`
}

// Field defines a field.
type Field struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Kind        FieldKind    `json:"kind"`
	Constraints []Constraint `json:"constraints"`
}

// IsValid determines whether or not the supplied value is valid.
func (f *Field) IsValid(v interface{}) []ValidationResult {
	res := make([]ValidationResult, len(f.Constraints))
	for i, c := range f.Constraints {
		ok, err := c.IsValid(v)
		res[i] = ValidationResult{
			Constraint: c.Name(),
			Valid:      ok,
		}

		if err != nil {
			res[i].Error = err.Error()
		}
	}

	return res
}

// UnmarshalJSON satisfies json.Unmarshaler and ensures that deserialization of
// Constraints works as expected.
func (f *Field) UnmarshalJSON(data []byte) error {
	var field field
	if err := json.Unmarshal(data, &field); err != nil {
		return err
	}

	f.Name = field.Name
	f.Description = field.Description
	f.Kind = field.Kind

	for _, c := range field.Constraints {
		cs, err := ParseConstraint(c)
		if err != nil {
			return err
		}

		f.Constraints = append(f.Constraints, cs)
	}

	return nil
}

// ValidationResult encapsulates the result for validating a field constraint.
type ValidationResult struct {
	Constraint string `json:"constraint"`
	Valid      bool   `json:"valid"`
	Error      string `json:"error"`
}
