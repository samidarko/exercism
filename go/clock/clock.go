package clock

import (
	"fmt"
	"strings"
)

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	return Clock{hour: hour, minute: minute}
}
func (c Clock) String() string {
	var output strings.Builder
	if c.hour < 10 {
		output.WriteString("0")
	}
	output.WriteString(fmt.Sprintf("%d", c.hour))
	output.WriteString(":")
	if c.minute < 10 {
		output.WriteString("0")
	}
	output.WriteString(fmt.Sprintf("%d", c.minute))
	return output.String()
}

func (c Clock) Add(t int) Clock {
	return Clock{hour: 0, minute: 0}
}

func (c Clock) Subtract(t int) Clock {
	return Clock{hour: 0, minute: 0}
}
