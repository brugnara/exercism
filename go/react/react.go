package react

type React struct{}

func New() React {
	return React{}
}

// CreateInput foo
func (r React) CreateInput(x int) InputCell {
	// return RInputCell{&x.(RCell)}
	return RInputCell{RCell{&x}}
}

// CreateCompute1 bar
func (r React) CreateCompute1(c Cell, fn func(int) int) ComputeCell {
	return RComputeCell{c, fn}
}

// CreateCompute2 baz
func (r React) CreateCompute2(c1, c2 Cell, fn func(int, int) int) ComputeCell {
	return RComputeCell2{c1, c2, fn}
}
