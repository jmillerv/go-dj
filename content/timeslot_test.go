package content

import (
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

func (t *testTime) Now() time.Time {
	now, _ := time.Parse(time.Kitchen, "11:27PM")
	return now
}

func TestTimes_IsScheduledNow(t1 *testing.T) {
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
				Begin: "11:00PM",
				End:   "11:59PM",
			},
			want: true,
		},
		{
			name: "Returns False",
			fields: fields{
				Begin: "11:28PM",
				End:   "10:47PM",
			},
			want: false,
		},
		{
			name: "Success: evaluates true for times that traverse days",
			fields: fields{
				Begin: "11:00PM",
				End:   "2:30AM",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			timeTest := &testTime{}
			t := &Timeslot{
				Current: timeTest.Now(),
				Begin:   tt.fields.Begin,
				End:     tt.fields.End,
			}
			assert.Equalf(t1, tt.want, t.IsScheduledNow(), "IsScheduledNow()")
		})
	}
}
