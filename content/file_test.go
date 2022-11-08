package content

import (
	"os"
	"testing"
)

func TestLocalFile_Get(t *testing.T) {
	type fields struct {
		Name    string
		Content *os.File
		Path    string
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
			l := &LocalFile{
				Name:    tt.fields.Name,
				Content: tt.fields.Content,
				Path:    tt.fields.Path,
			}
			if err := l.Get(); (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalFile_Play(t *testing.T) {
	type fields struct {
		Name    string
		Content *os.File
		Path    string
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
			l := &LocalFile{
				Name:    tt.fields.Name,
				Content: tt.fields.Content,
				Path:    tt.fields.Path,
			}
			if err := l.Play(); (err != nil) != tt.wantErr {
				t.Errorf("Play() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalFile_Stop(t *testing.T) {
	type fields struct {
		Name    string
		Content *os.File
		Path    string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LocalFile{
				Name:    tt.fields.Name,
				Content: tt.fields.Content,
				Path:    tt.fields.Path,
			}
			l.Stop()
		})
	}
}
