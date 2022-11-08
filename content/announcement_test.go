package content

import "testing"

func TestAnnouncement_Get(t *testing.T) {
	type fields struct {
		Name    string
		Content []byte
		URL     string
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
			a := &Announcement{
				Name:    tt.fields.Name,
				Content: tt.fields.Content,
				URL:     tt.fields.URL,
				Path:    tt.fields.Path,
			}
			if err := a.Get(); (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAnnouncement_Play(t *testing.T) {
	type fields struct {
		Name    string
		Content []byte
		URL     string
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
			a := &Announcement{
				Name:    tt.fields.Name,
				Content: tt.fields.Content,
				URL:     tt.fields.URL,
				Path:    tt.fields.Path,
			}
			if err := a.Play(); (err != nil) != tt.wantErr {
				t.Errorf("Play() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAnnouncement_Stop(t *testing.T) {
	type fields struct {
		Name    string
		Content []byte
		URL     string
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
			a := &Announcement{
				Name:    tt.fields.Name,
				Content: tt.fields.Content,
				URL:     tt.fields.URL,
				Path:    tt.fields.Path,
			}
			a.Stop()
		})
	}
}
