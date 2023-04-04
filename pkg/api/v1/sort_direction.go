package v1

// ToSQL returns either ASC or DESC depending on the value of d.
func (d SortDirection) ToSQL() string {
	if d == SortDirection_SORT_DIRECTION_DESCENDING {
		return "DESC"
	}

	return "ASC"
}
