package content

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
		{
			name: "Error: implement me",
			fields: fields{
				Name:    "Test",
				Content: nil,
				URL:     "test.com",
				Path:    "/path/to/file",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Announcement{
				Name:    tt.fields.Name,
				Content: tt.fields.Content,
				URL:     tt.fields.URL,
				Path:    tt.fields.Path,
			}
			err := a.Get()
			if err != nil && tt.wantErr {
				assert.Error(t, err)
				return
			}
		})
	}
}
