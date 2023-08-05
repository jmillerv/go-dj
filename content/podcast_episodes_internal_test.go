// nolint:TODO https://github.com/jmillerv/go-dj/issues/16
package content

// file labeled _internal_test because none of these functions are public.

import (
	"testing"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/assert"
)

func Test_contains(t *testing.T) {
	type args struct {
		guids       []string
		episodeGuid string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Success: Returns True",
			args: args{},
			want: true,
		},
		{
			name: "Success: Returns False",
			args: args{},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := contains(tt.args.guids, tt.args.episodeGuid)
			assert.ObjectsAreEqual(tt.want, got)
		})
	}
}

func Test_podcastCacheData_fromCache(t *testing.T) {
	type fields struct {
		Guids     []string
		TTY       string
		CacheDate time.Time
	}
	type args struct {
		cacheData any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *podcastCacheData
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &podcastCacheData{
				Guids:     tt.fields.Guids,
				TTY:       tt.fields.TTY,
				CacheDate: tt.fields.CacheDate,
			}
			assert.Equalf(t, tt.want, p.fromCache(tt.args.cacheData), "fromCache(%v)", tt.args.cacheData)
		})
	}
}

func Test_podcasts_getNewestEpisode(t *testing.T) {
	type fields struct {
		Episodes []*gofeed.Item
	}
	tests := []struct {
		name   string
		fields fields
		want   episode
	}{
		//nolint:godox // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &podcasts{
				Episodes: tt.fields.Episodes,
			}
			assert.Equalf(t, tt.want, p.getNewestEpisode(), "getNewestEpisode()")
		})
	}
}

func Test_podcasts_getOldestEpisode(t *testing.T) {
	type fields struct {
		Episodes []*gofeed.Item
	}
	tests := []struct {
		name   string
		fields fields
		want   episode
	}{
		//nolint:godox // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &podcasts{
				Episodes: tt.fields.Episodes,
			}
			assert.Equalf(t, tt.want, p.getOldestEpisode(), "getOldestEpisode()")
		})
	}
}

func Test_podcasts_getRandomEpisode(t *testing.T) {
	type fields struct {
		Episodes []*gofeed.Item
	}
	tests := []struct {
		name   string
		fields fields
		want   episode
	}{
		//nolint:godox // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &podcasts{
				Episodes: tt.fields.Episodes,
			}
			assert.Equalf(t, tt.want, p.getRandomEpisode(), "getRandomEpisode()")
		})
	}
}
