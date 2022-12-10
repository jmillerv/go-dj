package content

import (
	"github.com/mmcdole/gofeed"
	"math/rand"
	"time"
)

type podcasts struct {
	Episodes []*gofeed.Item
	// add podcast cache
}

func (p *podcasts) getNewestEpisode() *episode {
	var newestEpisode *episode
	for _, ep := range p.Episodes {
		var date *time.Time
		// TODO if played, log that it's in the cache, and skip to the next episode
		if ep.PublishedParsed.After(*date) {
			date = ep.PublishedParsed
			newestEpisode.Item = ep
			newestEpisode.EpURL = ep.Enclosures[0].URL
			newestEpisode.EpExtension = ep.Enclosures[0].Type
		}
	}
	return newestEpisode
}

func (p *podcasts) getOldestEpisode() *episode {
	var oldestEpisode *episode
	// TODO if played, log that it's in the cache, and skip to the next episode
	for _, ep := range p.Episodes {
		var date *time.Time
		if ep.PublishedParsed.Before(*date) {
			date = ep.PublishedParsed
			oldestEpisode.Item = ep
			oldestEpisode.EpURL = ep.Enclosures[0].URL
			oldestEpisode.EpExtension = ep.Enclosures[0].Type
		}
	}

	return oldestEpisode
}

func (p *podcasts) getRandomEpisode() *episode {
	var randomEpisode *episode
	rand.Seed(time.Now().UnixNano())
	item := p.Episodes[rand.Intn(len(p.Episodes))]
	// TODO if played, log that it's in the cache, and skip to the next episode
	randomEpisode.Item = item
	randomEpisode.EpExtension = item.Enclosures[0].Type
	randomEpisode.EpURL = item.Enclosures[0].URL
	return randomEpisode
}

func (p *podcasts) checkIfPlayed(guid string) bool {
	return false
}

type episode struct {
	Item        *gofeed.Item // Keep this to hold the additional data
	EpExtension string
	EpURL       string
}
