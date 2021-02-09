package react

import "log"

type React struct {
	cells map[Cell][]interface{}
}

func New() Reactor {
	r := &React{
		map[Cell][]interface{}{},
	}

	return r
}

func (r React) sync(n, o int, cell Cell) {
	for _, c := range r.cells[cell] {
		if cc, ok := c.(*rComputeCell2); ok {
			// cc.Callbacks()
			log.Println(n, o)
			if cc.ChangesWith(n, o, cell) {
				cc.Callbacks()
				r.sync(n, o, cc)
			}
			break
		}
		if cc, ok := c.(*rComputeCell); ok {
			//
			if cc.Computed(n) != cc.Computed(o) {
				cc.Callbacks()
				r.sync(cc.Computed(n), cc.Computed(o), cc)
				break
			}
		}
	}
}

// CreateInput foo
func (r *React) CreateInput(x int) (ret InputCell) {
	// return RInputCell{&x.(RCell)}
	ret = &RInputCell{RCell(x), func(n, o int, cell Cell) {
		log.Println("Callback with new:", n, "and old:", o, cell)
		r.sync(n, o, cell)
	}}
	return
}

// CreateCompute1 bar
func (r *React) CreateCompute1(c Cell, fn func(int) int) ComputeCell {
	ret := &rComputeCell{Cell: c, variator: fn}
	r.cells[c] = append(r.cells[c], ret)
	return ret
}

// CreateCompute2 baz
func (r React) CreateCompute2(c1, c2 Cell, fn func(int, int) int) ComputeCell {
	ret := &rComputeCell2{c1: c1, c2: c2, variator: fn}
	r.cells[c1] = append(r.cells[c1], ret)
	r.cells[c2] = append(r.cells[c2], ret)

	return ret
}
