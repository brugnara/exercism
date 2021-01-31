package react

type RCanceler struct {
	callback func()
}

func (c RCanceler) Cancel() {
	c.callback()
}
