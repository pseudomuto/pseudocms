package validation

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Constraint defines an interface for field constraints.
type Constraint interface {
	// IsValid determines whether the supplied value is valid. It returns the result
	// of the validatity test and any errors that may have occurred during the test.
	IsValid(interface{}) (bool, error)

	// String returns the serialized form of this constraint. This is used for JSON
	// and database serialization. The result of String can be passed to
	// ParseConstraint without expectation of error.
	String() string
}

// ParseConstraint parses the given string representation of a constraint into a
// Constraint object. The assumption here is that the string representation is
// equal to the value returned by calling Name on the result.
func ParseConstraint(c string) (Constraint, error) {
	parts := strings.SplitN(c, "(", 2)

	switch parts[0] {
	case "required":
		return IsRequired(), nil
	case "maxLength":
		v, err := strconv.Atoi(strings.SplitN(parts[1], ")", 2)[0])
		if err != nil {
			return nil, err
		}

		return MaxLength(v), nil
	case "minLength":
		v, err := strconv.Atoi(strings.SplitN(parts[1], ")", 2)[0])
		if err != nil {
			return nil, err
		}

		return MinLength(v), nil
	}

	return nil, fmt.Errorf("unknown constraint: %s", c)
}

type constraint struct {
	name string
}

func (c *constraint) Name() string { return c.name }

func (c *constraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.name)
}

func (c *constraint) String() string { return c.name }
