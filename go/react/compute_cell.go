package react

// RComputeCell fobar
type RComputeCell struct {
	Cell

	variator  func(int) int
	callbacks map[int]func(int)
}

// AddCallback foobar
func (r RComputeCell) AddCallback(fn func(int)) Canceler {
	return RCanceler(10)
}

// Value overrides RInputCell Value()
func (r RComputeCell) Value() int {
	return r.variator(r.Cell.Value())
}

// 2

type RComputeCell2 struct {
	c1 Cell
	c2 Cell

	variator func(int, int) int
}

func (r RComputeCell2) AddCallback(fn func(int)) Canceler {
	return RCanceler(10)
}

// func (r RComputeCell2) SetValue(i int) {}

func (r RComputeCell2) Value() int {
	return r.variator(r.c1.Value(), r.c2.Value())
}
