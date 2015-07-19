package clock

import "fmt"

// TestVersion is a test version.
const TestVersion = 2

const (
	minPerHour = 60
	hourPerDay = 24
)

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
	if minutes < 0 || minutes >= minPerHour {
		n := minutes / minPerHour
		hour += n
		minutes -= n * minPerHour

		if minutes < 0 {
			minutes += minPerHour
			hour--
		}
	}

	if hour < 0 || hour >= hourPerDay {
		hour %= hourPerDay

		if hour < 0 {
			hour += hourPerDay
		}
	}

	return Clock{
		hour: hour,
		min:  minutes,
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
	c.hour += minutes / minPerHour
	m := c.min + minutes%minPerHour
	if m >= minPerHour {
		c.hour++
		m -= minPerHour
	} else if m < 0 {
		c.hour--
		m += minPerHour
	}
	c.min = m

	if c.hour >= hourPerDay {
		c.hour %= hourPerDay
	} else if c.hour < 0 {
		c.hour += hourPerDay
	}

	return c
}
