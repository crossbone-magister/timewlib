package timewlib

import (
	"fmt"
	"time"
)

type Interval struct {
	start time.Time
	end   time.Time
	Tags  []string
}

func NewInterval(startHour int, startMinute int, endHour int, endMinute int) *Interval {
	now := time.Now()
	return &Interval{
		start: time.Date(now.Year(), now.Month(), now.Day(), startHour, startMinute, 0, 0, now.Local().Location()),
		end:   time.Date(now.Year(), now.Month(), now.Day(), endHour, endMinute, 0, 0, now.Local().Location()),
	}
}

func (i *Interval) StartDate() (int, time.Month, int) {
	return i.start.Date()
}

func (i *Interval) EndDate() (int, time.Month, int) {
	return i.end.Date()
}

func (i *Interval) StartHour() int {
	return i.start.Hour()
}

func (i *Interval) EndHour() int {
	return i.end.Hour()
}

func (i *Interval) StartMinute() int {
	return i.start.Minute()
}

func (i *Interval) EndMinute() int {
	return i.end.Minute()
}

func (i *Interval) Duration() time.Duration {
	return i.end.Sub(i.start)
}

func (i *Interval) IsSameHour() bool {
	return i.start.Hour() == i.end.Hour()
}

func (i Interval) String() string {
	return fmt.Sprintf("[%s-%s]", i.start, i.end)
}

func (i Interval) ShortString() string {
	return fmt.Sprintf("%s-%s", i.start.Format("1504"), i.end.Format("1504"))
}
