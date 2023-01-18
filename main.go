package main

import (
	"os"
	"time"

	"github.com/jmillerv/go-dj/cache"
	"github.com/jmillerv/go-dj/content"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"zgo.at/zcache"
)

const (
	configFile = "config.yml"
)

func main() {
	app := &cli.App{
		Name:    "Go DJ",
		Usage:   "Daemon that schedules audio programming content",
		Version: "0.0.1",
		Commands: cli.Commands{
			{
				Name:      "start",
				Aliases:   []string{"s"},
				Usage:     "start",
				UsageText: "starts the daemon from the config",
				Action: func(c *cli.Context) {
					log.Info("creating schedule from config")
					scheduler, err := content.NewScheduler(configFile)
					if err != nil {
						log.WithError(err).Error("unable to run go-dj")
					}
					ttl, err := time.ParseDuration(scheduler.Content.PlayedPodcastTTL)
					if err != nil {
						log.WithError(err).Error("unable to parse played podcast ttl")
					}
					// create cache
					cache.PodcastPlayedCache = zcache.New(ttl, ttl)

					// hydrate podcast
					content.HydratePodcastCache()

					if content.Shuffled {
						log.Info("playing shuffled content")
						err = scheduler.Shuffle()
						if err != nil {
							log.WithError(err).Error("unable to run go-dj")
						}
						return
					}
					// run content normally
					err = scheduler.Run()
					if err != nil {
						log.WithError(err).Error("unable to run go-dj")
					}
				},
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:        "random",
						Usage:       "Start your radio station w/ randomized schedule",
						Required:    false,
						Hidden:      false,
						Destination: &content.Shuffled,
					},
					cli.BoolFlag{
						Name:        "pod-oldest",
						Usage:       "podcasts will play starting with the oldest first",
						Required:    false,
						Hidden:      false,
						Destination: &content.PodcastPlayerOrderOldest,
					},
					cli.BoolFlag{
						Name:        "pod-random",
						Usage:       "podcasts will play in a random order",
						Required:    false,
						Hidden:      false,
						Destination: &content.PodcastPlayOrderRandom,
					},
				},
			},
			{
				Name:      "clear-cache",
				Aliases:   []string{"clear"},
				Usage:     "./go-dj clear-cache",
				UsageText: "deletes the in memory and locally saved podcast cache",
				Action: func(c *cli.Context) {
					log.Info("clearing cache")
					scheduler, err := content.NewScheduler(configFile)
					ttl, err := time.ParseDuration(scheduler.Content.PlayedPodcastTTL)
					if err != nil {
						log.WithError(err).Error("unable to parse played podcast ttl")
					}
					// create cache
					cache.PodcastPlayedCache = zcache.New(ttl, ttl)

					// hydrate podcast
					content.HydratePodcastCache()
					err = cache.ClearPodcastPlayedCache()
					if err != nil {
						log.WithError(err).Error("unable to clear podcast played cache")
					}
				},
			},
		},
		Author: "Jeremiah Miller",
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func init() {
	content.Shuffled = false
	content.PodcastPlayerOrderOldest = false
	content.PodcastPlayOrderRandom = false
}
