package cache

import "testing"

func TestClearPodcastPlayedCache(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ClearPodcastPlayedCache(); (err != nil) != tt.wantErr {
				t.Errorf("ClearPodcastPlayedCache() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
