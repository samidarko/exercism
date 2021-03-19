package clock

import (
	"fmt"
)

// Clock stores time in minutes for a day
type Clock struct {
	time int
}

const totalMinutes = 1440 // total minutes in a day

// New is a Clock constructor
func New(hour, minute int) Clock {
	time := hour*60 + minute
	if time < 0 {
		time = totalMinutes + (time - (time/totalMinutes)*totalMinutes)
	}
	if time >= totalMinutes {
		time = time - (time/totalMinutes)*totalMinutes
	}
	return Clock{time}
}

// String display time as "HH:MM"
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour(), c.Minute())
}

// Add minutes to Clock
func (c Clock) Add(minutes int) Clock {
	return New(c.Hour(), c.Minute()+minutes)
}

// Subtract minutes to Clock
func (c Clock) Subtract(minutes int) Clock {
	return New(c.Hour(), c.Minute()-minutes)
}

// Hour returns which hour of the day
func (c Clock) Hour() int {
	return c.time / 60
}

// Minute returns which minute of the day
func (c Clock) Minute() int {
	return c.time % 60
}
