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

//const (
//	Early      Timeslot = "early"
//	Morning    Timeslot = "morning"
//	Breakfast  Timeslot = "breakfast"
//	Midmorning Timeslot = "midmorning"
//	Afternoon  Timeslot = "afternoon"
//	Commute    Timeslot = "commute"
//	Evening    Timeslot = "evening"
//	Late       Timeslot = "late"
//	Overnight  Timeslot = "overnight"
//	All        Timeslot = "all"
//)
//
//type Timeslot string
//
//type Slot struct {
//	Begin string
//	End   string
//}

var Shuffled bool

//
//var TimeslotMap = map[Timeslot]*Slot{
//	Early:      {"4:00 AM", "6:00 AM"},
//	Morning:    {"6:00 AM", "8:00 AM"},
//	Breakfast:  {"8:00 AM", "11:00 AM"},
//	Midmorning: {"11:00 AM", "2:00 PM"},
//	Afternoon:  {"2:00 PM", "5:00 PM"},
//	Commute:    {"5:00 PM", "7:00 PM"},
//	Evening:    {"7:00 PM", "11:00 PM"},
//	Late:       {"11:00 PM", "2:00 AM"},
//	Overnight:  {"2:00 AM", "4:00 AM"},
//	All:        {"12:00 AM", "12:00 PM"},
//}

type Scheduler struct {
	Content struct {
		Programs []*Program
	}
}

func (s *Scheduler) Run() error {
	log.Info("Starting Daemon")

	log.Infof("Press ESC to quit")
	// set up the loop to continuously check for key entries

	// if randomized mode do x

	// setup signal listeners
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)
	exitchnl := make(chan int)

	// check content from scheduler and run through it
	for _, p := range s.Content.Programs {
		log.Debugf("program %v", formatter.StructToIndentedString(p))
		// Check Timeslots
		if p.Timeslot.IsScheduledNow() {
			log.Infof("getting media type: %v", p.Type)
			content := p.GetMedia()
			log.Debugf("media struct: %v", content)
			content.Get()
			go func() {
				for {
					stop := <-sigchnl
					s.Stop(stop, content)
				}
			}()
			err := content.Play()
			if err != nil {
				return err
			} // play will block until done
		}
	}
	// if radio station start 30 minute counter.
	// smartly allocate programs to timeslots based on length if known
	// if time between TimeSlotMap do x
	// play program from that slot.
	// wait for program to finish
	// get next
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
