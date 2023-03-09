package validation

import (
	"errors"
	"fmt"
)

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
