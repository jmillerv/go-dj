package content

import (
	"fmt"
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

var Shuffled bool

type Scheduler struct {
	Content struct {
		// Expire time for played podcast cache
		PlayedPodcastTTL string
		// Duration between the loop pausing and checking the content against the schedule.
		CheckInterval string
		Programs      []*Program
	}
}

func (s *Scheduler) Run() error {
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
			log.Debugf("program %v", formatter.StructToIndentedString(p))

			// if content is scheduled, retrieve and play
			scheduled := p.Timeslot.IsScheduledNow(now)
			if scheduled {
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
				if p.getMediaType() == webRadioContent {
					go func() {
						duration := getDurationToEndTime(p.Timeslot.End) // might cause an index out of range issue
						stopCountDown(content, duration, &wg)
					}()
					go func() {
						log.Info("playing web radio inside of a go routine")
						wg.Add(1)
						err = content.Play()
						if err != nil {
							log.WithError(err).Error("Run::content.Play")
						} // play will block until done
					}()
				} else if p.getMediaType() == podcastContent {
					go func() {
						podcast := content.(*Podcast)
						log.Infof("podcast player duration %s", podcast.Player.duration)
						stopCountDown(content, podcast.Player.duration, &wg)
					}()
					go func() {
						log.Info("playing podcast inside of a go routine")
						wg.Add(1)
						err = content.Play()
						if err != nil {
							log.WithError(err).Error("Run::content.Play")
						} // play will block until done
					}()
				} else {
					err = content.Play() // play will block until done
					if err != nil {
						return err
					}
				}
			}
			log.Info("paused while go routines are running")
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

// Shuffle plays through the config content at random
func (s *Scheduler) Shuffle() error {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s.Content.Programs),
		func(i, j int) {
			s.Content.Programs[i], s.Content.Programs[j] = s.Content.Programs[j], s.Content.Programs[i]
		})

	// setup signal listeners
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)
	exitchnl := make(chan int)

	for _, p := range s.Content.Programs {
		log.Debugf("program %v", formatter.StructToIndentedString(p))
		log.Infof("getting media type: %v", p.Type)
		content := p.GetMedia()
		log.Debugf("media struct: %v", content)
		err := content.Get()
		if err != nil {
			return err
		}
		go func() {
			for {
				stop := <-sigchnl
				s.Stop(stop, content)
			}
		}()
		err = content.Play()
		if err != nil {
			return err
		} // play will block until done
	}

	exitcode := <-exitchnl
	os.Exit(exitcode)
	return nil
}

func (s *Scheduler) Stop(signal os.Signal, media Media) {
	if signal == syscall.SIGTERM {
		log.Info("Got kill signal. ")
		if media != nil {
			media.Stop()
		}
		log.Info("Program will terminate now.")
		os.Exit(0)
	} else if signal == syscall.SIGINT {
		if media != nil {
			media.Stop()
		}
		log.Info("Got CTRL+C signal")
		if media != nil {
			media.Stop()
		}
		fmt.Println("Closing.")
		os.Exit(0)
	}
}

func (s *Scheduler) getNextProgram(index int) *Program {
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

// stopCountDown takes in a Media and duration and starts a ticker to stop the playing content
func stopCountDown(content Media, period time.Duration, wg *sync.WaitGroup) {
	log.Infof("remaining time playing this stream %v", period)
	t := time.NewTicker(period)
	defer t.Stop()
	for {
		select {
		case <-t.C: // call content.Stop
			log.Info("stopping content")
			err := content.Stop()
			if err != nil {
				log.WithError(err).Error("stopCountDown::error stopping content")
			}
			// typecast content as WebRadio
			webRadio, ok := content.(*WebRadio)
			if ok {
				// only send a wg.Done() signal if the web radio has stopped playing.
				if !webRadio.Player.isPlaying {
					wg.Done()
				}
			}
			// typecast content as Podcast
			podcast, ok := content.(*Podcast)
			if ok {
				if !podcast.Player.isPlaying {
					wg.Done()
				}
			}
			log.Info("content stopped")
			return
		}
	}
}

// getDurationToEndTime determines how much time in seconds needs to pass before the next program starts.
// TODO look at this function and timeslot.go's IsScheduleNow() and attempt to refactor to remove duplicate code.
func getDurationToEndTime(currentProgramEnd string) time.Duration {
	current := time.Now()
	// get date info for string
	date := time.Date(current.Year(), current.Month(), current.Day(), 0, 0, 0, 0, current.Location())
	year, month, day := date.Date()

	// convert ints to dateString
	dateString := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)

	// parse the date and the config time
	// parsed times are returned in 2022-12-05 15:05:00 +0000 UTC format which doesn't appear to have a const in the time package
	parsedProgramEnd, _ := dateparse.ParseAny(dateString + " " + currentProgramEnd)

	// matched parse time to fixed zone time
	currentProgramEndTime := time.Date(parsedProgramEnd.Year(), parsedProgramEnd.Month(), parsedProgramEnd.Day(), parsedProgramEnd.Hour(), parsedProgramEnd.Minute(), parsedProgramEnd.Second(), parsedProgramEnd.Nanosecond(), current.Location())

	duration := currentProgramEndTime.Sub(current)
	return duration
}
