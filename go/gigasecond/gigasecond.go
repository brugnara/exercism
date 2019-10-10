package gigasecond

import "time"

const gigaSecond = 1 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000

// AddGigasecond adds a gigaSecond (1 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000)
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(gigaSecond))
}
