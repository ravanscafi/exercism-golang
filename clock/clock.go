package clock

import "fmt"

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	return Clock{
		hour:   ((23 + (hour+minute/60)%24) + (60+minute%60)/60) % 24,
		minute: (60 + minute%60) % 60,
	}
}

func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", c.hour, c.minute)
}
