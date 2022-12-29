package content

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewScheduler(t *testing.T) {
	type args struct {
		file string
	}
	type want struct {
		scheduler *Scheduler
	}
	tests := []struct {
		name    string
		args    args
		want    *want
		wantErr bool
	}{
		{
			name: "Success: Returns scheduler",
			args: args{
				file: "../config.test.yml",
			},
			want: &want{
				scheduler: &Scheduler{
					Content: struct {
						PlayedPodcastTTL string
						CheckInterval    string
						Programs         []*Program
					}{
						PlayedPodcastTTL: "3h",
						CheckInterval:    "1m",
						Programs: []*Program{
							{
								Name:   "gettysburg10",
								Source: "./static/gettysburg10.wav",
								Timeslot: &Timeslot{
									Begin: "11:00PM",
									End:   "11:30PM",
								},
								Type: MediaType("file"),
							},
						}},
				},
			},
			wantErr: false,
		},
		{
			name: "Error: failed to read in config file",
			args: args{
				file: "./fakefile.yaml",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error: scheduler is empty",
			args: args{
				file: "../config.test-bad.yml",
			},
			want:    nil,
			wantErr: true,
		},
	}
	// TODO make test pass when running in parallel and troubleshoot race condition.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewScheduler(tt.args.file)
			if err != nil && tt.wantErr {
				assert.Error(t, err)
				return
			}
			if err == nil && tt.wantErr {
				log.Info("error is nil, expected error")
				t.Fail()
				return
			}
			if err != nil {
				log.WithError(err)
				t.Fail()
				return
			}
			assert.ObjectsAreEqual(tt.want.scheduler, got)
		})
	}
}

func Test_getDurationBetweenPrograms(t *testing.T) {
	type args struct {
		endTime string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success: Returns Correct Duration",
			args: args{
				endTime: "9:30PM",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Info(getDurationtoEndTime(tt.args.endTime))
		})
	}
}
