package content

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type TimeProvider interface {
	Now() time.Time
}

type testTime struct {
	TimeProvider
}

func (testTime *testTime) Now() time.Time {
	tz, _ := time.LoadLocation("EST")
	now := time.Date(2022, 12, 05, 23, 27, 0, 0, tz)
	log.Infof("testTime %v", now)
	return now
}

func TestTimes_IsScheduledNow(t *testing.T) {
	t.Parallel()
	type fields struct {
		Current time.Time
		Begin   string
		End     string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Returns True",
			fields: fields{
				Begin: "11:00 PM",
				End:   "11:59 PM",
			},
			want: true,
		},
		{
			name: "Returns False",
			fields: fields{
				Begin: "11:28 PM",
				End:   "10:47 PM",
			},
			want: false,
		},
		{
			name: "Success: evaluates true for times that traverse days",
			fields: fields{
				Begin: "11:00 PM",
				End:   "2:30 AM",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t1 *testing.T) {
			t1.Parallel()
			t := &Timeslot{
				Begin: tt.fields.Begin,
				End:   tt.fields.End,
			}
			assert.Equalf(t1, tt.want, t.IsScheduledNow((&testTime{}).Now()), "IsScheduledNow()")
		})
	}
}
