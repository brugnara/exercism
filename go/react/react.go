package react

type React struct {
	update    chan Cell
	done      chan bool
	cells     map[Cell][]*rComputeCell
	cells2    map[Cell][]*rComputeCell2
	cellValue map[Cell]int
}

func New() Reactor {
	chn := make(chan Cell)
	done := make(chan bool)
	r := &React{
		chn,
		done,
		map[Cell][]*rComputeCell{},
		map[Cell][]*rComputeCell2{},
		map[Cell]int{},
	}

	go func() {
		for c := range chn {
			// 1
			if _, ok := r.cells[c]; ok {
				for _, cell := range r.cells[c] {
					val := cell.Value()
					if val != r.cellValue[c] {
						r.cellValue[c] = val
						cell.Callbacks()
					}
				}
			}

			// 2
			if _, ok := r.cells2[c]; ok {
				for _, cell := range r.cells2[c] {
					val := cell.Value()
					if val != r.cellValue[c] {
						r.cellValue[c] = val
						cell.Callbacks()
					}
				}
			}

			//
			done <- true
		}
	}()
	return r
}

// CreateInput foo
func (r *React) CreateInput(x int) (ret InputCell) {
	// return RInputCell{&x.(RCell)}
	ret = &RInputCell{RCell(x), r.update, r.done}
	return
}

// CreateCompute1 bar
func (r React) CreateCompute1(c Cell, fn func(int) int) ComputeCell {
	ret := &rComputeCell{Cell: c, variator: fn}
	// r.cells = append(r.cells, ret)
	r.cells[c] = append(r.cells[c], ret)
	r.cellValue[c] = ret.Value()
	return ret
}

// CreateCompute2 baz
func (r React) CreateCompute2(c1, c2 Cell, fn func(int, int) int) ComputeCell {
	ret := &rComputeCell2{c1: c1, c2: c2, variator: fn}
	//
	r.cells2[c1] = append(r.cells2[c1], ret)
	r.cellValue[c1] = ret.Value()
	r.cells2[c2] = append(r.cells2[c2], ret)
	r.cellValue[c2] = ret.Value()
	//
	return ret
}
