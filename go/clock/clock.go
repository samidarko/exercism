package clock

import (
	"fmt"
)

// Clock stores time in minutes for a day
type Clock struct {
	time int
}

const (
	minutesByDay  = 1440 // total minutes in a day
	minutesByHour = 60   // total minutes in a day
)

// New is a Clock constructor
func New(hour, minute int) Clock {
	time := hour*minutesByHour + minute
	if time < 0 {
		time = minutesByDay + (time - (time/minutesByDay)*minutesByDay)
	}
	if time >= minutesByDay {
		time = time - (time/minutesByDay)*minutesByDay
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
	return c.time / minutesByHour
}

// Minute returns which minute of the day
func (c Clock) Minute() int {
	return c.time % minutesByHour
}
