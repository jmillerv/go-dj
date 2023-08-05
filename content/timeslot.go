package content

import (
	"strconv"
	"time"

	"github.com/araddon/dateparse"
	log "github.com/sirupsen/logrus"
)

// Times represents timeslots and are parsed in a 24hour format.
type Timeslot struct {
	Begin string
	End   string
}

// IsScheduledNow checks the current time and returns a bool if the time falls within the range.
func (t *Timeslot) IsScheduledNow(current time.Time) bool {
	// get date info for string
	date := time.Date(current.Year(), current.Month(), current.Day(), 0, 0, 0, 0, current.Location())
	year, month, day := date.Date()

	// convert ints to dateString
	dateString := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)

	// parse the date and the config time
	// parsed times are returned in 2022-12-05 15:05:00 +0000 UTC format
	// which doesn't appear to have a const in the time package
	parsedStartTime, _ := dateparse.ParseAny(dateString + " " + t.Begin)
	parsedEndTime, _ := dateparse.ParseAny(dateString + " " + t.End)

	// matched parse time to fixed zone time
	startTime := time.Date(parsedStartTime.Year(), parsedStartTime.Month(), parsedStartTime.Day(),
		parsedStartTime.Hour(), parsedStartTime.Minute(), parsedStartTime.Second(),
		parsedStartTime.Nanosecond(), current.Location())

	endTime := time.Date(parsedEndTime.Year(), parsedEndTime.Month(), parsedEndTime.Day(), parsedEndTime.Hour(),
		parsedEndTime.Minute(), parsedEndTime.Second(), parsedEndTime.Nanosecond(), current.Location())

	return inTimeSpan(startTime, endTime, current)
}

func inTimeSpan(start, end, current time.Time) bool {
	log.WithField("start", start).WithField("current", current).WithField("end", end).
		Info("timeslot::inTimeSpan: configured times")
	// handle scheduling that traverses days.
	tz, _ := time.LoadLocation("UTC")

	if end.Before(start) && current.After(start) {
		return true
	}

	return current.In(tz).After(start) && current.In(tz).Before(end)
}
