package content

import (
	"math/rand"
	"time"

	"github.com/jmillerv/go-dj/cache"
	"github.com/mmcdole/gofeed"
)

type podcasts struct {
	Episodes []*gofeed.Item
	// add podcast cache
}

type podcastCacheData struct {
	Guids     []string  `json:"guids"`
	TTY       string    `json:"tty"`
	CacheDate time.Time `json:"cacheDate"`
}

func (p *podcastCacheData) fromCache(cacheData any) *podcastCacheData {
	data, ok := cacheData.(podcastCacheData)
	if ok {
		return &data
	}

	return nil
}

//nolint:ineffassign,staticcheck,wastedassign
func (p *podcasts) getNewestEpisode() episode {
	var newestEpisode episode
	var date *time.Time //nolint:wsl // declarations are fine to cuddle

	date = p.Episodes[0].PublishedParsed

	for i, ep := range p.Episodes {
		// check for cacheData cache
		cacheData, cacheHit := cache.PodcastPlayedCache.Get(defaultPodcastCache)
		if cacheHit {
			retrieved := (&podcastCacheData{}).fromCache(cacheData)
			if contains(retrieved.Guids, ep.GUID) {
				continue
			}
		}

		date = p.Episodes[i].PublishedParsed // update date

		if ep.PublishedParsed.After(*date) || ep.PublishedParsed.Equal(*date) {
			date = ep.PublishedParsed
			newestEpisode.Item = ep
			newestEpisode.EpURL = ep.Enclosures[0].URL
			newestEpisode.EpExtension = ep.Enclosures[0].Type
		}
	}

	return newestEpisode
}

//nolint:ineffassign,staticcheck,wastedassign
func (p *podcasts) getOldestEpisode() episode {
	var oldestEpisode episode
	var date *time.Time //nolint:wsl //  it's fine to cuddle declarations

	date = p.Episodes[0].PublishedParsed // update date

	for i, ep := range p.Episodes {
		cacheData, cacheHit := cache.PodcastPlayedCache.Get(ep.GUID)
		if cacheHit {
			retrieved := (&podcastCacheData{}).fromCache(cacheData)
			if contains(retrieved.Guids, ep.GUID) {
				continue
			}
		}

		date = p.Episodes[i].PublishedParsed // update date
		if ep.PublishedParsed.Before(*date) || ep.PublishedParsed.Equal(*date) {
			date = ep.PublishedParsed
			oldestEpisode.Item = ep
			oldestEpisode.EpURL = ep.Enclosures[0].URL
			oldestEpisode.EpExtension = ep.Enclosures[0].Type
		}
	}

	return oldestEpisode
}

func (p *podcasts) getRandomEpisode() episode {
	var randomEpisode episode

	rand.Seed(time.Now().UnixNano())

	item := p.Episodes[rand.Intn(len(p.Episodes))]
	_, cacheHit := cache.PodcastPlayedCache.Get(item.GUID)

	// block until cacheHit != true
	for cacheHit {
		item = p.Episodes[rand.Intn(len(p.Episodes))]

		_, cacheHit = cache.PodcastPlayedCache.Get(item.GUID)
		if cacheHit {
			continue
		}
	}

	randomEpisode.Item = item
	randomEpisode.EpExtension = item.Enclosures[0].Type
	randomEpisode.EpURL = item.Enclosures[0].URL

	return randomEpisode
}

type episode struct {
	Item        *gofeed.Item // Keep this to hold the additional data
	EpExtension string
	EpURL       string
}

func contains(guids []string, episodeGuid string) bool {
	for _, v := range guids {
		if v == episodeGuid {
			return true
		}
	}

	return false
}
