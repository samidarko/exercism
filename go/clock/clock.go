package clock

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	return Clock{hour: hour, minute: minute}
}
func (c Clock) String() string {
	return ""
}

func (c Clock) Add(t int) Clock {
	return Clock{hour: 0, minute: 0}
}

func (c Clock) Subtract(t int) Clock {
	return Clock{hour: 0, minute: 0}
}
