package main

import (
	"github.com/jmillerv/go-dj/content"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

const (
	configFile = "config.dev.yml"
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
					s := content.NewScheduler(configFile)
					if content.Shuffled {
						log.Info("playing shuffled content")
						err := s.Shuffle()
						if err != nil {
							log.WithError(err).Error("unable to run go-dj")
						}
						return
					}
					// run content normally
					err := s.Run()
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
				},
			},
		},
		Author: "Jeremiah Miller",
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
