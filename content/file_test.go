// nolint:TODO https://github.com/jmillerv/go-dj/issues/16
package content

import (
	"io"
	"os"
	"testing"

	"github.com/faiface/beep"
	"github.com/stretchr/testify/assert"
)

const (
	envCI = "CI"
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
		//nolint:godox // TODO: Add test cases.
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
		//nolint:godox // TODO: Add test cases.
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
		//nolint:godox // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LocalFile{
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

func TestLocalFile_getEstimatedFileDuration(t *testing.T) {
	if os.Getenv(envCI) == "true" {
		t.Skipf("CI detected, skipping")
	}
	type fields struct {
		Name             string
		Content          *os.File
		Path             string
		decodeReader     func(r io.Reader) (s beep.StreamSeekCloser, format beep.Format, err error)
		decodeReadCloser func(rc io.ReadCloser) (s beep.StreamSeekCloser, format beep.Format, err error)
		fileType         string
	}
	mp3File := &LocalFile{
		Name:    "piano_six_seconds.mp3",
		Path:    "../static/piano_six_seconds.mp3",
		Content: nil,
	}
	wavFile := &LocalFile{
		Name:    "CantinaBand3.wav",
		Path:    "../static/CantinaBand3.wav",
		Content: nil,
	}
	flacFile := &LocalFile{
		Name:    "JosefSuk-Meditation.flac",
		Path:    "../static/JosefSuk-Meditation.flac",
		Content: nil,
	}
	oggFile := &LocalFile{
		Name:    "",
		Content: nil,
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "wav file",
			fields: fields{
				Name:     wavFile.Name,
				Content:  wavFile.Content,
				Path:     wavFile.Path,
				fileType: wavFile.fileType,
			},
			want: "3",
		},
		{
			name: "mp3 file",
			fields: fields{
				Name:     mp3File.Name,
				Content:  mp3File.Content,
				Path:     mp3File.Path,
				fileType: mp3File.fileType,
			},
			want: "",
		},
		{
			name: "ogg file",
			fields: fields{
				Name:     oggFile.Name,
				Content:  oggFile.Content,
				Path:     oggFile.Path,
				fileType: oggFile.fileType},
			want: "",
		},
		{
			name: "FLAC file",
			fields: fields{
				Name:     flacFile.Name,
				Content:  flacFile.Content,
				Path:     flacFile.Path,
				fileType: flacFile.fileType,
			},
			want: "",
		},
		{
			name: "default",
			fields: fields{
				fileType: "mp4",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LocalFile{
				Name:    tt.fields.Name,
				Content: tt.fields.Content,
				Path:    tt.fields.Path,
			}
			assert.Equalf(t, tt.want, l.getEstimatedFileDuration(), "getEstimatedFileDuration()")
		})
	}
}
