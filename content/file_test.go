package content

import (
	"github.com/faiface/beep"
	"os"
	"reflect"
	"testing"
)

func TestLocalFile_checkHeader(t *testing.T) {
	t.Parallel()
	type fields struct {
		Name    string
		Content *os.File
		Path    string
	}
	tests := []struct {
		name       string
		fields     fields
		wantS      beep.StreamSeekCloser
		wantFormat beep.Format
		wantErr    bool
	}{
		{
			name: "Error: returns default",
			fields: fields{
				Name:    "",
				Content: nil,
				Path:    "",
			},
			wantS:      nil,
			wantFormat: beep.Format{},
			wantErr:    false,
		},
		{
			name:       "Success: returns wav",
			fields:     fields{},
			wantS:      nil,
			wantFormat: beep.Format{},
			wantErr:    false,
		},
		{
			name:       "Success: returns mp3",
			fields:     fields{},
			wantS:      nil,
			wantFormat: beep.Format{},
			wantErr:    false,
		},
		{
			name:       "Success: returns flac",
			fields:     fields{},
			wantS:      nil,
			wantFormat: beep.Format{},
			wantErr:    false,
		},
		{
			name:       "Success: returns ogg",
			fields:     fields{},
			wantS:      nil,
			wantFormat: beep.Format{},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := &LocalFile{
				Name:    tt.fields.Name,
				Content: tt.fields.Content,
				Path:    tt.fields.Path,
			}
			gotS, gotFormat, err := l.checkHeader()
			if (err != nil) != tt.wantErr {
				t.Errorf("checkHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("checkHeader() gotS = %v, want %v", gotS, tt.wantS)
			}
			if !reflect.DeepEqual(gotFormat, tt.wantFormat) {
				t.Errorf("checkHeader() gotFormat = %v, want %v", gotFormat, tt.wantFormat)
			}
		})
	}
}
