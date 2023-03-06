package content

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Constraint defines an interface for field constraints.
type Constraint interface {
	// Name returns the unique name for this contraint.
	Name() string

	// IsValid determines whether the supplied value is valid. It returns the result
	// of the validatity test and any errors that may have occurred during the test.
	IsValid(interface{}) (bool, error)
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

// IsRequired returns a constraint which requires the value to not be empty or 0
// depending on the type.
func IsRequired() Constraint {
	return &required{constraint{name: "required"}}
}

// MaxLength returns a constraint which ensures the given value is at most n in
// length.
func MaxLength(n int) Constraint {
	return &maxLength{
		constraint: constraint{
			name: fmt.Sprintf("maxLength(%d)", n),
		},
		length: n,
	}
}

// MinLength returns a constraint which ensures the given value is at least n in
// length.
func MinLength(n int) Constraint {
	return &minLength{
		constraint: constraint{
			name: fmt.Sprintf("minLength(%d)", n),
		},
		length: n,
	}
}

type constraint struct {
	name string
}

func (c *constraint) Name() string { return c.name }

func (c *constraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.name)
}

type maxLength struct {
	constraint
	length int
}

func (m *maxLength) IsValid(v interface{}) (bool, error) {
	if val, ok := v.(string); ok {
		return len(val) <= m.length, nil
	}

	return false, errors.New("invalid type for maxLength")
}

type minLength struct {
	constraint
	length int
}

func (m *minLength) IsValid(v interface{}) (bool, error) {
	if val, ok := v.(string); ok {
		return len(val) >= m.length, nil
	}

	return false, errors.New("invalid type for minLength")
}

type required struct {
	constraint
}

func (r *required) IsValid(v interface{}) (bool, error) {
	if v == nil {
		return false, nil
	}

	switch val := v.(type) {
	case string:
		return val != "", nil
	case int:
		return val != 0, nil
	case uint:
		return val != uint(0), nil
	case int32:
		return val != int32(0), nil
	case uint32:
		return val != uint32(0), nil
	case int64:
		return val != int64(0), nil
	case uint64:
		return val != 0, nil
	case float32:
		return val != float32(0), nil
	case float64:
		return val != float64(0), nil
	}

	return true, nil
}
