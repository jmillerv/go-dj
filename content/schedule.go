package content

import (
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/araddon/dateparse"
	"github.com/jmillerv/go-utilities/formatter"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//nolint:gochecknoglobals // the globals here help but a refactor would be considered.
var Shuffled bool

type Scheduler struct {
	Content struct {
		// Expire time for played podcast cache
		PlayedPodcastTTL string
		// Duration between the loop pausing and checking the content against the schedule.
		CheckInterval string
		// Determines if go-dj will all a podcast to finish playing or force the next scheduled program
		StrictPodcastSchedule bool
		Programs              []*Program
	}
}

func (s *Scheduler) Run() error { //nolint:godox,funlen,gocognit,cyclop,nolintlint // TODO: consider refactoring
	var wg sync.WaitGroup

	log.Info("Starting Daemon")

	// setup signal listeners
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)

	exitchnl := make(chan int)

	totalPrograms := len(s.Content.Programs)
	programIndex := 0

	// run operation in loop
	for programIndex <= totalPrograms {
		// check content from scheduler and run through it
		// for loop that can be forced to continue from a go routine
		for _, p := range s.Content.Programs {
			now := time.Now()

			log.Infof("program %v", formatter.StructToIndentedString(p))

			// if content is scheduled, retrieve and play
			scheduled := p.Timeslot.IsScheduledNow(now)
			if scheduled { //nolint:nestif // consider refactoring
				log.Infof("scheduler.Run::getting media type: %v", p.Type)

				content := p.GetMedia()

				log.Debugf("media struct: %v", content)

				err := content.Get() // retrieve contents from file
				if err != nil {
					return err
				}

				// setup channel for os.Exit signal
				go func() {
					for {
						stop := <-sigchnl
						s.Stop(stop, content)
					}
				}()

				// if p.getMediaType is webRadioContent or podcastContent start a timer and stop content from inside a go routine
				// because these are streams rather than files they behave differently from local content.
				switch p.getMediaType() { //nolint:exhaustive,dupl // TODO consider refactoring into function
				case webRadioContent:
					log.Info("web radio content detected")

					webRadio := content.(*WebRadio)                          //nolint:forcetypeassert // TODO: type checking
					webRadio.Duration = getDurationToEndTime(p.Timeslot.End) // might cause an index out of range issue

					log.Debug("scheduler.Run::add 1 to waitgroup")
					wg.Add(1)

					go func() {
						defer wg.Done()
						log.Info("Run::playing web radio inside of a go routine")

						err := content.Play()
						if err != nil {
							log.WithError(err).Error("Run::content.Play")
						} // play will block until done
					}()

				case podcastContent:
					log.Info("podcast content detected")

					podcast := content.(*Podcast) //nolint:forcetypeassert // TODO: type checking

					// If the StrictPodcastSchedule is set to false, use the duration of the podcast to set the countdown.
					if !s.Content.StrictPodcastSchedule {
						podcast.Duration = podcast.Player.duration
					} else {
						podcast.Duration = getDurationToEndTime(p.Timeslot.End)
					}

					wg.Add(1)

					go func() {
						defer wg.Done()
						log.Info("playing podcast inside of a go routine")

						err = content.Play()
						if err != nil {
							log.WithError(err).Error("Run::content.Play")
						} // play will block until done
					}()

				default:
					err = content.Play() // play will block until done
					if err != nil {
						return err
					}
				}
			}

			log.Info("scheduler paused while go routines are running")
			wg.Wait() // pause

			if !p.Timeslot.IsScheduledNow(now) {
				log.WithField("IsScheduledNow", p.Timeslot.IsScheduledNow(now)).
					WithField("current time", time.Now().
						Format(time.Kitchen)).Infof("media not scheduled")
			}

			programIndex++ // increment index

			// check programs for scheduled content at regular interval
			if programIndex > totalPrograms {
				programIndex = 0

				// get the scheduled check interval from the scheduler
				interval, err := time.ParseDuration(s.Content.CheckInterval)
				if err != nil {
					return err
				}

				go func() {
					for {
						stop := <-sigchnl
						s.Stop(stop, nil) // passing nil because there is no media to stop.
					}
				}()

				// pause the loop
				log.WithField("pause interval", s.Content.CheckInterval).Info("loop paused, will resume after pause interval")
				time.Sleep(interval)
			}
		}
	}

	exitcode := <-exitchnl
	os.Exit(exitcode)

	return nil
}

// Shuffle plays through the config content at random.
func (s *Scheduler) Shuffle() error { //nolint:godox,funlen,gocognit,cyclop,nolintlint // TODO: consider refactoring
	wg := new(sync.WaitGroup)

	log.Info("Starting Daemon")

	rand.Shuffle(len(s.Content.Programs),
		func(i, j int) {
			s.Content.Programs[i], s.Content.Programs[j] = s.Content.Programs[j], s.Content.Programs[i]
		})

	// setup signal listeners
	sigchnl := make(chan os.Signal, 1)

	signal.Notify(sigchnl)

	exitchnl := make(chan int)

	totalPrograms := len(s.Content.Programs)
	programIndex := 0

	// run operation in loop
	for programIndex <= totalPrograms {
		for _, p := range s.Content.Programs {
			log.Debugf("program %v", formatter.StructToIndentedString(p))
			log.Infof("getting media type: %v", p.Type)
			content := p.GetMedia()
			log.Debugf("media struct: %v", content)

			err := content.Get()
			if err != nil {
				return err
			}

			// setup channel for os.Exit signal
			go func() {
				for {
					stop := <-sigchnl
					s.Stop(stop, content)
				}
			}()

			// if p.getMediaType is webRadioContent or podcastContent start a timer and stop content from inside a go routine
			// because these are streams rather than files they behave differently from local content.
			//nolint:godox,nolintlint // TODO: consider refactoring switch into single function used in Run() & Shuffle()
			switch p.getMediaType() { //nolint:dupl,exhaustive
			case webRadioContent:
				log.Info("web radio content detected")

				webRadio := content.(*WebRadio)                          //nolint:forcetypeassert // TODO: type checking
				webRadio.Duration = getDurationToEndTime(p.Timeslot.End) // might cause an index out of range issue

				log.Debug("scheduler.Run::add 1 to waitgroup")
				wg.Add(1)

				go func() {
					defer wg.Done()
					log.Info("Run::playing web radio inside of a go routine")

					err := content.Play()
					if err != nil {
						log.WithError(err).Error("Run::content.Play")
					} // play will block until done
				}()
			case podcastContent:
				log.Info("podcast content detected")

				podcast := content.(*Podcast) //nolint:forcetypeassert // TODO: type checking

				// If the StrictPodcastSchedule is set to false, use the duration of the podcast to set the countdown.
				if !s.Content.StrictPodcastSchedule {
					podcast.Duration = podcast.Player.duration
				} else {
					podcast.Duration = getDurationToEndTime(p.Timeslot.End)
				}

				wg.Add(1)

				go func() {
					defer wg.Done()
					log.Info("playing podcast inside of a go routine")

					err = content.Play()
					if err != nil {
						log.WithError(err).Error("Run::content.Play")
					} // play will block until done
				}()
			default:
				err = content.Play() // play will block until done
				if err != nil {
					return err
				}
			}

			log.Info("paused while go routines are running")

			wg.Wait() // pause

			programIndex++ // increment index
		}
	}

	exitcode := <-exitchnl
	os.Exit(exitcode)

	return nil
}

func (s *Scheduler) Stop(signal os.Signal, media Media) {
	if signal == syscall.SIGTERM { //nolint:nestif // TODO: consider refactoring
		log.Info("Got kill signal. ")

		if media != nil {
			err := media.Stop()
			if err != nil {
				log.WithError(err).Error("scheduler.Stop::error stopping media")
			}
		}

		log.Info("Program will terminate now.")

		os.Exit(0)
	} else if signal == syscall.SIGINT {
		if media != nil {
			err := media.Stop()
			if err != nil {
				log.WithError(err).Error("scheduler.Stop::error stopping media")
			}
		}

		log.Info("Got CTRL+C signal")

		if media != nil {
			media.Stop() //nolint:errcheck
		}

		log.Println("Closing.")

		os.Exit(0)
	}
}

func (s *Scheduler) getNextProgram(index int) *Program { //nolint:unused
	return s.Content.Programs[index]
}

func NewScheduler(file string) (*Scheduler, error) {
	log.Info("Loading Config File from: ", file)
	viper.SetConfigType("yaml")
	viper.SetConfigFile(file)
	viper.SetDefault("CheckInterval", "10m")     // default Check Interval
	viper.SetDefault("PlayedPodcastTTL", "730h") // default Cache TTL 730h is ~1 month

	if err := viper.ReadInConfig(); err != nil {
		log.WithField("file", file).WithError(err).Error("Failed to read in config file")

		return nil, err
	}

	scheduler := new(Scheduler)

	if err := viper.Unmarshal(scheduler); err != nil {
		log.WithError(err).Error("unable to unmarshal config into struct")

		return nil, err
	}

	if scheduler.Content.Programs == nil {
		return nil, errors.New("scheduler is empty")
	}

	log.Info("config loaded", formatter.StructToIndentedString(scheduler))

	return scheduler, nil
}

// getDurationToEndTime determines how much time in seconds needs to pass before the next program starts.
// TODO: examine function and timeslot.go's IsScheduleNow(), attempt to refactor to remove duplicate code.
//
//nolint:godox //ignore the below line
func getDurationToEndTime(currentProgramEnd string) time.Duration {
	current := time.Now()
	// get date info for string
	date := time.Date(current.Year(), current.Month(), current.Day(), 0, 0, 0, 0, current.Location())
	year, month, day := date.Date()

	// convert ints to dateString
	dateString := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)

	// parse the date and the config time
	// parsed times are returned in 2022-12-05 15:05:00 +0000 UTC format
	// which doesn't appear to have a const in the time package
	parsedProgramEnd, _ := dateparse.ParseAny(dateString + " " + currentProgramEnd)

	// matched parse time to fixed zone time
	currentProgramEndTime := time.Date(parsedProgramEnd.Year(), parsedProgramEnd.Month(), parsedProgramEnd.Day(),
		parsedProgramEnd.Hour(), parsedProgramEnd.Minute(), parsedProgramEnd.Second(),
		parsedProgramEnd.Nanosecond(), current.Location())

	duration := currentProgramEndTime.Sub(current)

	return duration
}
