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
	PlayOrder PlayOrder
}

type PlayOrder string

// Get parses a podcast feed and sets the most recent episode as the Podcast content.
func (p *Podcast) Get() error {
	var ep episode
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(p.URL)
	if err != nil {
		return err
	}
	// traverse links
	for _, item := range feed.Items {
		pods.Episodes = append(pods.Episodes, item)
	}

	switch p.PlayOrder {
	case playOrderNewest:
		ep = pods.getNewestEpisode()
		break
	case playOrderOldest:
		ep = pods.getOldestEpisode()
	case playOrderRandom:
		ep = pods.getRandomEpisode()
	}

	// setup podcast stream
	podcastStream.playerName = streamPlayerName
	podcastStream.url = ep.EpURL
	podcastStream.command = exec.Command(podcastStream.playerName, "-quiet", podcastStream.url)
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
	log.Infof("streaming from %v ", p.URL)
	if !p.Player.isPlaying {
		err := p.Player.command.Start()
		if err != nil {
			return errors.Wrap(err, "error starting podcast streamPlayer")
		}
		p.Player.isPlaying = true
		done := make(chan bool)
		func() {
			p.Player.pipeChan <- p.Player.out
			done <- true
		}()
		<-done
	}
	return nil
}

func (p *Podcast) Stop() error {
	log.Infof("Stopping stream from %v ", p.URL)
	if p.Player.isPlaying {
		p.Player.isPlaying = false
		_, err := p.Player.in.Write([]byte("q"))
		if err != nil {
			log.WithError(err).Error("error stopping web radio streamPlayerName: w.Player.in.Write()")
		}
		err = p.Player.in.Close()
		if err != nil {
			log.WithError(err).Error("error stopping web radio streamPlayerName: w.Player.in.Close()")
		}
		err = p.Player.out.Close()
		if err != nil {
			log.WithError(err).Error("error stopping web radio streamPlayerName: w.Player.out.Close()")
		}
		p.Player.command = nil

		p.Player.url = ""
	}
	return nil
}
