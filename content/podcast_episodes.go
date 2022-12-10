package content

import (
	"github.com/mmcdole/gofeed"
)

type podcasts struct {
	Episodes []*gofeed.Item
	// add podcast cache
}

func (p *podcasts) getNewestEpisode() *episode {
	return &episode{}
}

func (p *podcasts) getOldestEpisode() *episode {

	// get most recent by data

	// check most recent against cache
	// if played, get next most recent
	return &episode{
		Item:        nil,
		EpExtention: "",
		EpURL:       "",
	}
}

func (p *podcasts) getRandomEpisode() *episode {
	return &episode{}
}

func (p *podcasts) checkIfPlayed() bool {
	return false
}

type episode struct {
	Item        *gofeed.Item // Keep this to hold the additional data
	EpExtention string
	EpURL       string
}
