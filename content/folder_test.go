package content

import (
	"os"
	"reflect"
	"testing"
)

func TestFolder_Get(t *testing.T) {
	type fields struct {
		Name    string
		Content []os.DirEntry
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
			f := &Folder{
				Name:    tt.fields.Name,
				Content: tt.fields.Content,
				Path:    tt.fields.Path,
			}
			if err := f.Get(); (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFolder_Play(t *testing.T) {
	type fields struct {
		Name    string
		Content []os.DirEntry
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
			f := &Folder{
				Name:    tt.fields.Name,
				Content: tt.fields.Content,
				Path:    tt.fields.Path,
			}
			if err := f.Play(); (err != nil) != tt.wantErr {
				t.Errorf("Play() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFolder_Stop(t *testing.T) {
	type fields struct {
		Name    string
		Content []os.DirEntry
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
			f := &Folder{
				Name:    tt.fields.Name,
				Content: tt.fields.Content,
				Path:    tt.fields.Path,
			}
			f.Stop()
		})
	}
}

func TestFolder_getLocalFile(t *testing.T) {
	type fields struct {
		Name    string
		Content []os.DirEntry
		Path    string
	}
	type args struct {
		file os.DirEntry
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *LocalFile
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Folder{
				Name:    tt.fields.Name,
				Content: tt.fields.Content,
				Path:    tt.fields.Path,
			}
			if got := f.getLocalFile(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLocalFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
