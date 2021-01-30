package react

// React is a reactor implementation
type React struct{}

// New returns a Reactor
func New() Reactor {
	return React{}
}

// 1, genera la "cella"
func (r React) CreateInput(x int) InputCell {
	return RInputCell{&x}
}

func (r React) CreateCompute1(c Cell, fn func(int) int) (o ComputeCell) {
	o = RComputeCell1{&c, fn}
	r.cells[c] = o
	//
	return
}

func (r React) CreateCompute2(c1, c2 Cell, fn func(int, int) int) ComputeCell {
	return RComputeCell2{&c1, &c2, fn}
}

//
type RInputCell struct {
	value *int
}

func (c RInputCell) SetValue(x int) {
	*c.value = x
}

func (c RInputCell) Value() int {
	return *c.value
}

//

type RComputeCell1 struct {
	cell1    *Cell
	variator func(int) int
}

func (c RComputeCell1) AddCallback(fn func(int)) Canceler {
	return RCanceler(10)
}

func (c RComputeCell1) Value() int {
	return c.variator((*c.cell1).Value())
}

//

type RComputeCell2 struct {
	c1, c2 *Cell

	variator func(int, int) int
}

func (c RComputeCell2) AddCallback(fn func(int)) Canceler {
	return RCanceler(10)
}

func (c RComputeCell2) Value() int {
	return c.variator(
		(*c.c1).Value(),
		(*c.c2).Value(),
	)
}

//

type RCanceler int

func (c RCanceler) Cancel() {}
