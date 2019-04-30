package clock

import (
	"fmt"
)

// Clock stores the hour and min of a clock
type Clock struct {
	min int
}

// New returns a Clock with the hour and min
func New(hour, min int) Clock {
	return Clock{min: correctClock(60*hour + min)}
}

// Add adds minutes to the clock
func (c Clock) Add(min int) Clock {
	return Clock{min: correctClock(c.min + min)}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.min/60.0, c.min%60)
}

//Subtract takes away the given minutes from the clock
func (c Clock) Subtract(min int) Clock {
	return Clock{min: correctClock(c.min - min)}
}

func correctClock(min int) int {
	min = min % 1440
	if min < 0 {
		min += 1440
	}
	return min
}
