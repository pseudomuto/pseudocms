package validation

// Result defines the result of a constraint validation.
type Result struct {
	Constraint Constraint `json:"constraint"`
	Valid      bool       `json:"valid"`
	Error      string     `json:"error"`
}

// Validate deterines whether or not v is valid based on the supplied constraints
func Validate(v interface{}, constraints ...Constraint) []Result {
	res := make([]Result, len(constraints))
	for i, c := range constraints {
		ok, err := c.IsValid(v)
		res[i] = Result{
			Constraint: c,
			Valid:      ok,
		}

		if err != nil {
			res[i].Error = err.Error()
		}
	}

	return res
}
