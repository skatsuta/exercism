package gigasecond

import "time"

// TestVersion is a test version.
const TestVersion = 2

// Birthday is my birthday.
var Birthday = time.Date(1987, 12, 18, 0, 0, 0, 0, time.Local)

// Gigasecond is 10^9 seconds.
const Gigasecond = 1e9

// AddGigasecond returns a Time added with 1 Gs.
func AddGigasecond(in time.Time) time.Time {
	return in.Add(time.Duration(Gigasecond * time.Second))
}
