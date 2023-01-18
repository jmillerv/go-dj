package content

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/jmillerv/go-dj/cache"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"zgo.at/zcache"
)

const (
	playOrderNewest            PlayOrder = "newest"
	playOrderOldest            PlayOrder = "oldest"
	playOrderRandom            PlayOrder = "random"
	defaultPodcastCache                  = "podcastCache"
	podcastCacheLocalFile                = "./cache/podcastCache.json"
	localFileTTY                         = "72h"
	defaultPodcastPlayDuration           = "1h"
)

var pods podcasts // holds the feed data for podcasts
var podcastStream streamPlayer
var podcastCache podcastCacheData

type Podcast struct {
	Name        string
	URL         string
	Player      streamPlayer
	PlayOrder   PlayOrder
	EpisodeGuid string
	TTL         time.Duration // cache expiration time
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
		break
	case playOrderRandom:
		ep = pods.getRandomEpisode()
		break
	}
	// set guid for cache when played
	if ep.Item != nil {
		p.EpisodeGuid = ep.Item.GUID
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

	// get podcast duration
	if ep.EpURL != "" {
		podcastStream.duration, err = time.ParseDuration(fmt.Sprintf("%ss", ep.Item.ITunesExt.Duration))
		if err != nil {
			log.Infof("error parsing duration, setting default duration")
			podcastStream.setDuration(defaultPodcastPlayDuration)
			log.WithError(err).Errorf("error parsing duration %s", ep.Item.ITunesExt.Duration)
		}
	} else {
		log.Infof("podcast lacks duration, setting default duration")
		podcastStream.setDuration(defaultPodcastPlayDuration)
		if err != nil {
			return errors.Wrap(err, "error parsing duration")
		}
	}

	// set isPlaying to false
	podcastStream.isPlaying = false

	p.Player = podcastStream

	return nil
}

// Play sends the audio to the output. It caches a played episode in the cache ofr later checks.
func (p *Podcast) Play() error {
	log.Infof("streaming from %v ", p.URL)
	if !p.Player.isPlaying {
		log.WithField("episode", p.EpisodeGuid).Info("setting podcast played cache")
		cacheData, cacheHit := cache.PodcastPlayedCache.Get(defaultPodcastCache)
		if cacheHit {
			podcastCache = cacheData.(podcastCacheData)
		}
		if p.EpisodeGuid != "" {
			podcastCache.Guids = append(podcastCache.Guids, p.EpisodeGuid)
		}
		err := p.setCache(&podcastCache)
		if err != nil {
			return err
		}
		err = p.Player.command.Start()
		if err != nil {
			return errors.Wrap(err, "podcast.Play::error starting podcast streamPlayer")
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
	log.Infof("poadcast.Stop::Stopping stream from %v ", p.URL)
	if p.Player.isPlaying {
		p.Player.isPlaying = false
		_, err := p.Player.in.Write([]byte("q"))
		if err != nil {
			log.WithError(err).Errorf("podcast.Stop::error stopping %s: w.Player.in.Write()", p.Player.playerName)
		}
		err = p.Player.in.Close()
		if err != nil {
			log.WithError(err).Errorf("podcast.Stop::error stopping %s: w.Player.in.Write()", p.Player.playerName)
		}
		err = p.Player.out.Close()
		if err != nil {
			log.WithError(err).Errorf("podcast.Stop::error stopping %s: w.Player.in.Write()", p.Player.playerName)
		}
		p.Player.command = nil

		p.Player.url = ""
	}
	return nil
}

// setCache updates the in memory cache and persists a local file
func (p *Podcast) setCache(cacheData *podcastCacheData) error {
	cache.PodcastPlayedCache.Set(defaultPodcastCache, cacheData, zcache.DefaultExpiration)
	cacheData.TTY = localFileTTY
	cacheData.CacheDate = time.Now() // This will keep the cache constantly refreshing every time an episode is played. // TODO improve solution
	file, err := json.MarshalIndent(cacheData, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(podcastCacheLocalFile, file, 0644)
	if err != nil {
		return err
	}
	return nil
}

// HydratePodcastCache populates the default podcast cache with a local file
func HydratePodcastCache() {
	// check if file exists
	file, err := os.ReadFile(podcastCacheLocalFile)
	if errors.Is(err, os.ErrNotExist) {
		// if file does not exist do not hydrate the cache
		return
	}
	data := podcastCacheData{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.WithError(err).Error("HydratePodcastCache::failed to unmarshal podcast cache local file")
		return
	}
	// check that TTY is within range of cacheDate
	duration, err := time.ParseDuration(data.TTY)
	if err != nil {
		log.WithError(err).Error("HydratePodcastCache::failed to parse tty")
		return
	}
	if !data.CacheDate.Before(data.CacheDate.Add(duration)) {
		// if TTY is not within range, do not hydrate
		return
	}
	cache.PodcastPlayedCache.Set(defaultPodcastCache, data, zcache.DefaultExpiration)
}
