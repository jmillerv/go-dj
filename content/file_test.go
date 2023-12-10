// nolint:TODO https://github.com/jmillerv/go-dj/issues/16
package content

import (
	"io"
	"os"
	"testing"

	"github.com/faiface/beep"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const (
	envCI = "CI"
)

func TestLocalFile_Get(t *testing.T) {
	t.Parallel()
	type fields struct {
		Name             string
		Content          *os.File
		Path             string
		decodeReader     func(r io.Reader) (s beep.StreamSeekCloser, format beep.Format, err error)
		decodeReadCloser func(rc io.ReadCloser) (s beep.StreamSeekCloser, format beep.Format, err error)
		fileType         string
	}
	type want struct {
		err error
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(t *testing.T, f *fields) *LocalFile
		wantErr bool
		want    want
	}{
		{
			name: "Success: Open file without error",
			fields: fields{
				Name:             "localfile_test",
				Path:             "./",
				decodeReader:     nil,
				decodeReadCloser: nil,
				fileType:         ".txt",
			},
			prepare: func(t *testing.T, f *fields) *LocalFile {
				t.Helper()
				return &LocalFile{
					Name:             f.Name,
					Path:             f.Path,
					decodeReader:     f.decodeReader,
					decodeReadCloser: f.decodeReadCloser,
					fileType:         f.fileType,
				}
			},
			wantErr: false,
		},
		{
			name: "Error: failed to open file",
			fields: fields{
				Name:             "adsadsada",
				Path:             "32190da",
				decodeReader:     nil,
				decodeReadCloser: nil,
				fileType:         ".txt",
			},
			prepare: func(t *testing.T, f *fields) *LocalFile {
				t.Helper()
				f.decodeReader = func(r io.Reader) (s beep.StreamSeekCloser, format beep.Format, err error) {
					return nil, beep.Format{}, err
				}
				f.decodeReadCloser = func(rc io.ReadCloser) (s beep.StreamSeekCloser, format beep.Format, err error) {
					return nil, beep.Format{}, err
				}
				return &LocalFile{
					Name:             f.Name,
					Path:             f.Path,
					decodeReader:     f.decodeReader,
					decodeReadCloser: f.decodeReadCloser,
					fileType:         f.fileType,
				}
			},
			wantErr: true,
			want: want{
				err: errors.New("unable to open file from path open 32190da: no such file or directory"),
			},
		},
	}
	for _, tt := range tests {
		tt := tt // pin
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// create temporary file
			tmpFile, err := os.CreateTemp("", "localfile_test.txt")
			if err != nil {
				t.Fatalf("failed to create temporary file: %v", err)
			}
			defer os.Remove(tmpFile.Name())

			getErr := tt.prepare(t, &tt.fields).Get()
			if tt.wantErr {
				if getErr == nil {
					t.Fatalf("getErr shoold not be nil")
				}
				assert.EqualError(t, getErr, tt.want.err.Error())
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
		{
			name: "mp3 file",
			fields: fields{
				File: &LocalFile{
					Name:     "piano_six_seconds.mp3",
					Path:     "../static/piano_six_seconds.mp3",
					Content:  nil,
					fileType: mp3File,
				},
			},
			want: "6.37",
		},
		{
			name: "ogg file",
			fields: fields{
				File: &LocalFile{
					Name:     "Example.ogg",
					Path:     "../static/Example.ogg",
					fileType: oggFile,
				},
			},
			want: "6.10",
		},
		{
			name: "FLAC file",
			fields: fields{
				File: &LocalFile{
					Name:     "JosefSuk-Meditation.flac",
					Path:     "../static/JosefSuk-Meditation.flac",
					Content:  nil,
					fileType: flacFile,
				},
			},
			want: "405.45",
		},
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
