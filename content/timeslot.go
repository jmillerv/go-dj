package content

import (
	"time"
)

// Times represents timeslots and are parsed in a 24hour format
type Timeslot struct {
	Current time.Time
	Begin   string
	End     string
}

// IsScheduledNow checks the current time and returns a bool if the time falls within the range
func (t *Timeslot) IsScheduledNow() bool {
	startTime, _ := time.Parse(time.Kitchen, t.Begin)
	endTime, _ := time.Parse(time.Kitchen, t.End)

	return inTimeSpan(startTime, endTime, t.Current)

}

func inTimeSpan(start, end, current time.Time) bool {

	// handle scheduling that traverses days.
	if end.Before(start) && current.After(start) {
		return true
	}

	return current.After(start) && current.Before(end)
}
