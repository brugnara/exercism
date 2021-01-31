package react

// RCell contains a cell value
type RCell int

// Value returns the cell value
func (c *RCell) Value() int {
	return int(*c)
}
