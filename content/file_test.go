// nolint:TODO https://github.com/jmillerv/go-dj/issues/16
package content_test

import (
	"io"
	"os"
	"testing"

	"github.com/faiface/beep"
	"github.com/jmillerv/go-dj/content"
)

func TestLocalFile_Get(t *testing.T) {
	type fields struct {
		Name             string
		Content          *os.File
		Path             string
		decodeReader     func(r io.Reader) (s beep.StreamSeekCloser, format beep.Format, err error)
		decodeReadCloser func(rc io.ReadCloser) (s beep.StreamSeekCloser, format beep.Format, err error)
		fileType         string
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
			l := &content.LocalFile{
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
		Name             string
		Content          *os.File
		Path             string
		decodeReader     func(r io.Reader) (s beep.StreamSeekCloser, format beep.Format, err error)
		decodeReadCloser func(rc io.ReadCloser) (s beep.StreamSeekCloser, format beep.Format, err error)
		fileType         string
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
			l := &content.LocalFile{
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
		Name             string
		Content          *os.File
		Path             string
		decodeReader     func(r io.Reader) (s beep.StreamSeekCloser, format beep.Format, err error)
		decodeReadCloser func(rc io.ReadCloser) (s beep.StreamSeekCloser, format beep.Format, err error)
		fileType         string
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
			l := &content.LocalFile{
				Name:    tt.fields.Name,
				Content: tt.fields.Content,
				Path:    tt.fields.Path,
			}
			if err := l.Stop(); (err != nil) != tt.wantErr {
				t.Errorf("Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
