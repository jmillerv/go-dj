package content_test

import (
	. "github.com/jmillerv/go-dj/content"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewScheduler(t *testing.T) {
	t.Parallel()
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
					Content: struct{ Programs []*Program }{Programs: []*Program{
						{
							Name:     "gettysburg10",
							Source:   "./static/gettysburg10.wav",
							Timeslot: Timeslot("afternoon"),
							Type:     MediaType("file"),
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
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
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
