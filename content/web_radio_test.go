// nolint:TODO https://github.com/jmillerv/go-dj/issues/16
package content_test

import (
	"testing"

	. "github.com/jmillerv/go-dj/content"
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
		//nolint:godox // TODO: Add test cases.
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
		//nolint:godox // TODO: Add test cases.
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
		//nolint:godox // TODO: Add test cases.
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
