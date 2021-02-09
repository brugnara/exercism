package react

// RInputCell struct
type RInputCell struct {
	RCell
	callback func(int, int, Cell)
}

// SetValue sets a value
func (ic *RInputCell) SetValue(i int) {
	oldVal := ic.Value()
	ic.RCell = RCell(i)
	if ic.callback != nil {
		ic.callback(i, oldVal, ic)
	}
}
