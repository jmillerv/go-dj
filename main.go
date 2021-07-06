package main

import (
	"github.com/jmillerv/go-dj/content"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

const (
	configFile = "config.yml"
)

func main() {
	app := &cli.App{
		Name:    "Go DJ",
		Usage:   "Daemon that schedules audio programming content",
		Version: "1.0.0",
		Commands: cli.Commands{
			{
				Name:      "Start",
				Aliases:   []string{"s"},
				Usage:     "start",
				UsageText: "starts the daemon from the config",
				Action: func() {
					s := content.NewScheduler(configFile)
					err := s.Run()
					if err != nil {
						log.WithError(err).Error("unable to run go-dj")
						return
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
