package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func New(hours, minutes int) Clock {
	h := hours + minutes/60
	m := minutes % 60

	if m < 0 {
		m += 60
		h -= 1
	}

	h = h % 24
	if h < 0 {
		h += 24
	}

	c := new(Clock)
	c.hour = h
	c.minute = m
	return *c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02v:%02v", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}
