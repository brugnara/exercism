package react

// RInputCell struct
type RInputCell struct {
	RCell
}

// SetValue sets a value
func (ic RInputCell) SetValue(i int) {
	*ic.value = i
}
