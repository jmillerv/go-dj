package content_test

import (
	. "github.com/jmillerv/go-dj/content"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProgram_GetMedia(t *testing.T) {
	t.Parallel()
	type fields struct {
		program *Program
	}
	tests := []struct {
		name   string
		fields fields
		want   Media
	}{
		{
			name: "Success: returns folder",
			fields: fields{
				program: &Program{
					Name:   "David Rovics Folder",
					Source: "./static/david_rovics",
					Timeslot: &Timeslot{
						Begin: "11:00PM",
						End:   "11:30PM",
					},
					Type: "folder",
				},
			},
			want: (&Program{
				Name:   "David Rovics Folder",
				Source: "./static/david_rovics",
				Timeslot: &Timeslot{
					Begin: "11:00PM",
					End:   "11:30PM",
				},
				Type: "folder",
			}).GetMedia(),
		},
		{
			name: "Success: returns file",
			fields: fields{
				program: &Program{
					Name:   "Piano Six Seconds",
					Source: "./static/piano_six_seconds.mp3",
					Timeslot: &Timeslot{
						Begin: "11:00PM",
						End:   "11:30PM",
					},
					Type: "file",
				},
			},
			want: (&Program{
				Name:   "Piano Six Seconds",
				Source: "./static/piano_six_seconds.mp3",
				Timeslot: &Timeslot{
					Begin: "11:00PM",
					End:   "11:30PM",
				},
				Type: "file",
			}).GetMedia(),
		},
		{
			name: "Success: returns podcast",
			fields: fields{
				program: &Program{
					Name:   "Tech Won't Save Us",
					Source: "https://feeds.buzzsprout.com/1004689.rss",
					Timeslot: &Timeslot{
						Begin: "11:00PM",
						End:   "11:30PM",
					},
					Type: "podcast",
				},
			},
			want: nil,
		},
		{
			name: "Success: returns web radio",
			fields: fields{
				program: &Program{
					Name:   "Indie Pop Rocks",
					Source: "https://somafm.com/indiepop.pls",
					Timeslot: &Timeslot{
						Begin: "11:00PM",
						End:   "11:30PM",
					},
					Type: "web_radio",
				},
			},
			want: (&Program{
				Name:   "Indie Pop Rocks",
				Source: "https://somafm.com/indiepop.pls",
				Timeslot: &Timeslot{
					Begin: "11:00PM",
					End:   "11:30PM",
				},
				Type: "web_radio",
			}).GetMedia(),
		},
		{
			name: "Returns nil",
			fields: fields{
				program: &Program{Type: "test"},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := &Program{
				Name:     tt.fields.program.Name,
				Source:   tt.fields.program.Source,
				Timeslot: tt.fields.program.Timeslot,
				Type:     tt.fields.program.Type,
			}
			got := p.GetMedia()
			assert.ObjectsAreEqual(tt.want, got)
		})
	}
}
