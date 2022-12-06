package content

import (
	"fmt"
	"github.com/jmillerv/go-utilities/formatter"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Shuffled bool

type Scheduler struct {
	Content struct {
		CheckInterval string
		Programs      []*Program
	}
}

func (s *Scheduler) Run() error {
	log.Info("Starting Daemon")

	log.Infof("Press ESC to quit")

	// setup signal listeners
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)
	exitchnl := make(chan int)

	totalPrograms := len(s.Content.Programs)
	programIndex := 0

	// run operation in loop
	for programIndex <= totalPrograms {
		// check content from scheduler and run through it
		for _, p := range s.Content.Programs {
			now := time.Now()
			log.Debugf("program %v", formatter.StructToIndentedString(p))

			if p.Timeslot.IsScheduledNow(now) {
				log.Infof("getting media type: %v", p.Type)
				content := p.GetMedia()
				log.Debugf("media struct: %v", content)
				err := content.Get() // retrieve contents from file
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

			if !p.Timeslot.IsScheduledNow(now) {
				log.WithField("IsScheduledNow", p.Timeslot.IsScheduledNow(now)).
					WithField("current time", time.Now().
						Format(time.Kitchen)).Infof("media not scheduled")
			}
			programIndex++ // increment index
			if programIndex > totalPrograms {
				programIndex = 0

				// get the scheduled check interval from the scheduler
				interval, err := time.ParseDuration(s.Content.CheckInterval)
				if err != nil {
					return err
				}
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
		media.Stop()
		log.Info("Program will terminate now.")
		os.Exit(0)
	} else if signal == syscall.SIGINT {
		media.Stop()
		log.Info("Got CTRL+C signal")
		media.Stop()
		fmt.Println("Closing.")
		os.Exit(0)
	}
}

//func getTimeSlot(t *time.Time) Timeslot {
//	// if t between certain times return Timeslot
//	return All
//}

func NewScheduler(file string) (*Scheduler, error) {
	log.Info("Loading Config File from: ", file)
	viper.SetConfigType("yaml")
	viper.SetConfigFile(file)
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
