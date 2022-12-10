package content

import (
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

const (
	playOrderNewest PlayOrder = "newest"
	playOrderOldest PlayOrder = "oldest"
	playOrderRandom PlayOrder = "random"
)

var pods podcasts // holds the feed data for podcasts
var podcastStream streamPlayer

type Podcast struct {
	Name      string
	URL       string
	Player    streamPlayer
	PlayOrder PlayOrder // options: newest, oldest, random
}

type PlayOrder string

// Get parses a podcast feed and sets the most recent episode as the Podcast content.
func (p *Podcast) Get() error {
	var ep *episode
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(p.URL)
	if err != nil {
		return err
	}
	log.Infof("test %v", feed.Description)
	// traverse links
	for _, item := range feed.Items {
		pods.Episodes = append(pods.Episodes, item)
	}

	switch p.PlayOrder {
	case playOrderNewest:
		ep = pods.getNewestEpisode()
	case playOrderOldest:
		ep = pods.getOldestEpisode()
	case playOrderRandom:
		ep = pods.getRandomEpisode()
	}

	// setup podcast stream
	podcastStream.playerName = ep.EpExtension
	podcastStream.url = ep.EpURL
	podcastStream.command = exec.Command(podcastStream.playerName, "-quiet", "-playlist", podcastStream.url)
	podcastStream.in, err = podcastStream.command.StdinPipe()
	if err != nil {
		return errors.Wrap(err, "error creating standard pipe in")
	}
	podcastStream.out, err = podcastStream.command.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, "error creating standard pipe out")
	}

	podcastStream.isPlaying = false

	p.Player = podcastStream

	return nil
}

// Play sends the audio to the output. It caches a played episode in the cache ofr later checks.
func (p *Podcast) Play() error {
	// play file
	// cache played episode
	panic("implement me")
}

func (p *Podcast) Stop() error {
	log.Infof("Stopping stream from %v ", p.URL)
	return nil
}
