package main

//nolint:gci
import (
	"io"
	"os"
	"time"

	"github.com/jmillerv/go-dj/helpers"

	"github.com/jmillerv/go-dj/cache"
	"github.com/jmillerv/go-dj/content"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"zgo.at/zcache"
)

const (
	configFile     = "config.yml"
	configOverride = "GODJ_CONFIG_OVERRIDE"
	logFile        = "/tmp/godj.log"
	logPermissions = 0o666
)

func main() { //nolint:funlen,cyclop,gocognit // main function can be longer & more complex.
	app := &cli.App{ //nolint:exhaustivestruct,exhaustruct
		Name:    "Go DJ",
		Usage:   "Daemon that schedules audio programming content",
		Version: "0.0.1",
		Commands: cli.Commands{
			{ //nolint:exhaustruct
				Name:      "start",
				Aliases:   []string{"s"},
				Usage:     "start",
				UsageText: "starts the daemon from the config",
				Action: func(c *cli.Context) { //nolint:revive
					var config string
					log.Info("creating schedule from config")
					if os.Getenv(configOverride) != "" {
						config = os.Getenv(configOverride)
					} else {
						config = configFile
					}

					scheduler, err := content.NewScheduler(config)
					if err != nil {
						log.WithError(err).Error("content.NewScheduler::unable to run go-dj")
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
							log.WithError(err).Error("scheduler.Shuffle::unable to run go-dj")
						}

						return
					}
					// run content normally
					err = scheduler.Run()
					if err != nil {
						log.WithError(err).Error("scheduler.Run::unable to run go-dj")
					}
				},
				Flags: []cli.Flag{
					cli.BoolFlag{ //nolint:exhaustivestruct,exhaustruct
						Name:        "random",
						Usage:       "Start your radio station w/ randomized schedule",
						Required:    false,
						Hidden:      false,
						Destination: &content.Shuffled,
					},
					cli.BoolFlag{ //nolint:exhaustivestruct,exhaustruct
						Name:        "pod-oldest",
						Usage:       "podcasts will play starting with the oldest first",
						Required:    false,
						Hidden:      false,
						Destination: &content.PodcastPlayerOrderOldest,
					},
					cli.BoolFlag{ //nolint:exhaustivestruct,exhaustruct
						Name:        "pod-random",
						Usage:       "podcasts will play in a random order",
						Required:    false,
						Hidden:      false,
						Destination: &content.PodcastPlayOrderRandom,
					},
				},
			},
			{ //nolint:exhaustruct
				Name:      "clear-cache",
				Aliases:   []string{"clear"},
				Usage:     "./go-dj clear-cache",
				UsageText: "deletes the in memory and locally saved podcast cache",
				Action: func(c *cli.Context) { //nolint:revive
					log.Info("clearing cache")

					// set the config
					var config string
					if os.Getenv(configOverride) != "" {
						config = os.Getenv(configOverride)
					} else {
						config = configFile
					}
					scheduler, err := content.NewScheduler(config)
					if err != nil {
						log.WithError(err).Error("content.NewScheduler::unable to create scheduler from config file")
					}
					ttl, err := time.ParseDuration(scheduler.Content.PlayedPodcastTTL)
					if err != nil {
						log.WithError(err).Error("unable to parse played podcast ttl")
					}
					// create cache
					cache.PodcastPlayedCache = zcache.New(ttl, ttl)

					err = cache.ClearPodcastPlayedCache()
					if err != nil {
						log.WithError(err).Error("unable to clear podcast played cache")
					}
				},
			},
			{ //nolint:exhaustruct
				Name:      "install-dependencies",
				Aliases:   []string{"deps"},
				Usage:     "./go-dj install-dependencies",
				UsageText: "installs necessary dependencies to run go-dj",
				Action: func(c *cli.Context) { //nolint:revive
					packages := []string{"libasound2-dev", "libudev-dev", "pkg-config"}
					missingPackages := []string{}

					for _, pkg := range packages {
						if !helpers.PackageIsInstalled(pkg) {
							missingPackages = append(missingPackages, pkg)
						}
					}

					if len(missingPackages) > 0 {
						log.Info("The following packages are missing:")
						for _, pkg := range missingPackages {
							log.Info("-", pkg)
						}

						log.Info("Installing missing packages...")
						for _, pkg := range missingPackages {
							err := helpers.InstallPackage(pkg)
							if err != nil {
								log.WithError(err).Error("Error installing package:")

								return
							}
							log.Infof("Package %s installed successfully.\n", pkg)
						}
					} else {
						log.Info("All required packages are already installed.")
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

// init sets global variables.
func init() {
	content.Shuffled = false
	content.PodcastPlayerOrderOldest = false
	content.PodcastPlayOrderRandom = false

	initLogger()
}

// initLogger creates the multiwriter, determines the log format for each destination, and sets the logfile location.
// at a later stage, it may be desirable to have different formats for standard out vs the log file.
// An example of how to do that can be found here https://github.com/sirupsen/logrus/issues/784#issuecomment-403765306
func initLogger() {
	// create a new file for logs
	logs, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, logPermissions)
	if err != nil {
		log.WithError(err).Error("unable to open log file")
	}

	// open the multiwriter
	multiWrite := io.MultiWriter(os.Stdout, logs)

	log.SetFormatter(&log.TextFormatter{ //nolint:exhaustruct // don't need this full enumerated
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})

	log.SetOutput(multiWrite)
}
