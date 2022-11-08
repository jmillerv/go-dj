package content

import (
	"reflect"
	"testing"
)

func TestProgram_GetMedia(t *testing.T) {
	type fields struct {
		Name     string
		Source   string
		Timeslot Timeslot
		Type     MediaType
	}
	tests := []struct {
		name   string
		fields fields
		want   Media
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				Name:     tt.fields.Name,
				Source:   tt.fields.Source,
				Timeslot: tt.fields.Timeslot,
				Type:     tt.fields.Type,
			}
			if got := p.GetMedia(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMedia() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgram_mediaFactory(t *testing.T) {
	type fields struct {
		Name     string
		Source   string
		Timeslot Timeslot
		Type     MediaType
	}
	tests := []struct {
		name   string
		fields fields
		want   Media
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				Name:     tt.fields.Name,
				Source:   tt.fields.Source,
				Timeslot: tt.fields.Timeslot,
				Type:     tt.fields.Type,
			}
			if got := p.mediaFactory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mediaFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}