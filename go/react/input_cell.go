package react

// RInputCell struct
type RInputCell struct {
	RCell
	update chan Cell
	done   chan bool
}

// SetValue sets a value
func (ic *RInputCell) SetValue(i int) {
	ic.RCell = RCell(i)
	ic.update <- ic
	<-ic.done
}
