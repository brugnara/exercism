package react

// RCell contains a cell value
type RCell struct {
	value *int
}

// Value returns the cell value
func (c RCell) Value() int {
	return *c.value
}
