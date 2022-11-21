package content_test

import (
	. "github.com/jmillerv/go-dj/content"
	"testing"
)

func TestWebRadio_Get(t *testing.T) {
	type fields struct {
		Name string
		URL  string
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
			w := &WebRadio{
				Name: tt.fields.Name,
				URL:  tt.fields.URL,
			}
			if err := w.Get(); (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebRadio_Play(t *testing.T) {
	type fields struct {
		Name string
		URL  string
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
			w := &WebRadio{
				Name: tt.fields.Name,
				URL:  tt.fields.URL,
			}
			if err := w.Play(); (err != nil) != tt.wantErr {
				t.Errorf("Play() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebRadio_Stop(t *testing.T) {
	type fields struct {
		Name string
		URL  string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WebRadio{
				Name: tt.fields.Name,
				URL:  tt.fields.URL,
			}
			w.Stop()
		})
	}
}
