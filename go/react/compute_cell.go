package react

// rComputeCell fobar
type rComputeCell struct {
	Cell

	variator  func(int) int
	callbacks map[int]func(int)
}

func (r rComputeCell) Callbacks() {
	for _, v := range r.callbacks {
		v(r.Value())
	}
}

// AddCallback foobar
func (r *rComputeCell) AddCallback(fn func(int)) (ret Canceler) {
	if r.callbacks == nil {
		r.callbacks = make(map[int]func(int))
	}
	index := len(r.callbacks)
	r.callbacks[index] = fn
	ret = &RCanceler{func() {
		delete(r.callbacks, index)
	}}
	return
}

// Value overrides RInputCell Value()
func (r rComputeCell) Value() int {
	return r.variator(r.Cell.Value())
}

// 2

type rComputeCell2 struct {
	c1 Cell
	c2 Cell

	variator  func(int, int) int
	callbacks map[int]func(int)
}

func (r rComputeCell2) Callbacks() {
	for _, v := range r.callbacks {
		v(r.Value())
	}
}

func (r rComputeCell2) AddCallback(fn func(int)) Canceler {
	if r.callbacks == nil {
		r.callbacks = make(map[int]func(int))
	}
	index := len(r.callbacks)
	r.callbacks[index] = fn
	return &RCanceler{func() {
		delete(r.callbacks, index)
	}}
}

// func (r RComputeCell2) SetValue(i int) {}

func (r rComputeCell2) Value() int {
	return r.variator(r.c1.Value(), r.c2.Value())
}
