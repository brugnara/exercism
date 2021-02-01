package clock

import "fmt"

// Clock stores a clock
type Clock struct {
	minutes int
}

// New returns a Clock instance
func New(h, m int) Clock {
	return Clock{Clock.Compute(Clock{}, h*60+m)}
}
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours(), c.min())
}

// Add adds the given minutes
func (c Clock) Add(m int) Clock {
	return Clock{c.Compute(c.minutes + m)}
}

// Subtract subtracts the given minutes
func (c Clock) Subtract(m int) Clock {
	return Clock{c.Compute(c.minutes - m)}
}

// Compute fixes given input returning values always between
// 00:00 and 23:59
func (c Clock) Compute(m int) int {
	for m < 0 {
		m += c.day()
	}
	return m % c.day()
}

func (c Clock) day() int {
	return 24 * 60
}

func (c Clock) hours() int {
	return c.minutes / 60
}

func (c Clock) min() int {
	return c.minutes % 60
}
