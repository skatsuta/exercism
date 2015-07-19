package clock

import "fmt"

// TestVersion is a test version.
const TestVersion = 2

// Clock handles times without dates.
//
// A Clock is able to add and subtract minutes to it.
//
// Two clocks that represent the same time are equal to each other.
type Clock struct {
	hour int
	min  int
}

// Time creates a Clock holding hours and minutes.
func Time(hour, minutes int) Clock {
	d, m := minutes/60, minutes%60
	if minutes < 0 {
		d -= 1
		m += 60
	}
	h := hour%24 + d
	if h < 0 {
		h += 24
	}
	return Clock{
		hour: h,
		min:  m,
	}
}

// String returns a string representing time such as "08:00".
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}

// Add returns a Clock added by minutes.
// Add also handle subtraction by accepting negative values.
// Values of Clock type work with the == operator.
func (c Clock) Add(minutes int) Clock {
	s := c.min + minutes
	d, m := s/60, s%60
	return Time(c.hour+d, m)
}
