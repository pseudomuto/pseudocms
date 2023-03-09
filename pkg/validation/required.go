package validation

// IsRequired returns a constraint which requires the value to not be empty or 0
// depending on the type.
func IsRequired() Constraint {
	return &required{constraint{name: "required"}}
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
