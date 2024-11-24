// nolint:TODO https://github.com/jmillerv/go-dj/issues/16
package content

import (
	"io"
	"os"
	"testing"

	"github.com/gopxl/beep/v2"
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
		File             *LocalFile
		decodeReader     func(r io.Reader) (s beep.StreamSeekCloser, format beep.Format, err error)
		decodeReadCloser func(rc io.ReadCloser) (s beep.StreamSeekCloser, format beep.Format, err error)
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "wav file",
			fields: fields{
				File: &LocalFile{
					Name:     "CantinaBand3.wav",
					Path:     "../static/CantinaBand3.wav",
					fileType: wavFile,
				},
			},
			want: "library doesn't support estimating wav files.",
		},
		// Uncomment and add appropriate files for local testing
		//{
		//	name: "mp3 file",
		//	fields: fields{
		//		File: &LocalFile{
		//			Name:     "piano_six_seconds.mp3",
		//			Path:     "../static/piano_six_seconds.mp3",
		//			Content:  nil,
		//			fileType: mp3File,
		//		},
		//	},
		//	want: "6.37",
		//},
		//{
		//	name: "ogg file",
		//	fields: fields{
		//		File: &LocalFile{
		//			Name:     "Example.ogg",
		//			Path:     "../static/Example.ogg",
		//			fileType: oggFile,
		//		},
		//	},
		//	want: "6.10",
		//},
		//{
		//	name: "FLAC file",
		//	fields: fields{
		//		File: &LocalFile{
		//			Name:     "JosefSuk-Meditation.flac",
		//			Path:     "../static/JosefSuk-Meditation.flac",
		//			Content:  nil,
		//			fileType: flacFile,
		//		},
		//	},
		//	want: "405.45",
		//},
		{
			name: "default",
			fields: fields{
				File: &LocalFile{fileType: "mp4"},
			},
			want: "unknown file type: can't determine duration",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LocalFile{
				Name:     tt.fields.File.Name,
				Path:     tt.fields.File.Path,
				fileType: tt.fields.File.fileType,
			}

			assert.Equalf(t, tt.want, l.getEstimatedFileDuration(), "getEstimatedFileDuration()")
		})
	}
}
