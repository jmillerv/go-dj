package content

import "testing"

func TestPodcast_Get(t *testing.T) {
	type fields struct {
		Name    string
		URL     string
		Path    string
		Content []byte
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
			p := &Podcast{
				Name:    tt.fields.Name,
				URL:     tt.fields.URL,
				Path:    tt.fields.Path,
				Content: tt.fields.Content,
			}
			if err := p.Get(); (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPodcast_Play(t *testing.T) {
	type fields struct {
		Name    string
		URL     string
		Path    string
		Content []byte
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
			p := &Podcast{
				Name:    tt.fields.Name,
				URL:     tt.fields.URL,
				Path:    tt.fields.Path,
				Content: tt.fields.Content,
			}
			if err := p.Play(); (err != nil) != tt.wantErr {
				t.Errorf("Play() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPodcast_Stop(t *testing.T) {
	type fields struct {
		Name    string
		URL     string
		Path    string
		Content []byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Podcast{
				Name:    tt.fields.Name,
				URL:     tt.fields.URL,
				Path:    tt.fields.Path,
				Content: tt.fields.Content,
			}
			p.Stop()
		})
	}
}