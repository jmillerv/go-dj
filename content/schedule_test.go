package content_test

import (
	. "github.com/jmillerv/go-dj/content"
	"os"
	"reflect"
	"testing"
)

func TestNewScheduler(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want *Scheduler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScheduler(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScheduler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheduler_Run(t *testing.T) {
	type fields struct {
		Content struct {
			Programs []*Program
		}
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scheduler{
				Content: tt.fields.Content,
			}
			if err := s.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScheduler_Shuffle(t *testing.T) {
	type fields struct {
		Content struct {
			Programs []*Program
		}
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scheduler{
				Content: tt.fields.Content,
			}
			if err := s.Shuffle(); (err != nil) != tt.wantErr {
				t.Errorf("Shuffle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScheduler_Stop(t *testing.T) {
	type fields struct {
		Content struct {
			Programs []*Program
		}
	}
	type args struct {
		signal os.Signal
		media  Media
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scheduler{
				Content: tt.fields.Content,
			}
			s.Stop(tt.args.signal, tt.args.media)
		})
	}
}
