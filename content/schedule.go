package content

import (
	"github.com/jmillerv/go-utilities/formatter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

const (
	Early      Timeslot = "early"
	Morning    Timeslot = "morning"
	Breakfast  Timeslot = "breakfast"
	Midmorning Timeslot = "midmorning"
	Afternoon  Timeslot = "afternoon"
	Commute    Timeslot = "commute"
	Evening    Timeslot = "evening"
	Late       Timeslot = "late"
	Overnight  Timeslot = "overnight"
	All        Timeslot = "all"
)

type Timeslot string

type Slot struct {
	Begin string
	End   string
}

var Shuffled bool

var TimeslotMap = map[Timeslot]*Slot{
	Early:      {"4:00 AM", "6:00 AM"},
	Morning:    {"6:00 AM", "8:00 AM"},
	Breakfast:  {"8:00 AM", "11:00 AM"},
	Midmorning: {"11:00 AM", "2:00 PM"},
	Afternoon:  {"2:00 PM", "5:00 PM"},
	Commute:    {"5:00 PM", "7:00 PM"},
	Evening:    {"7:00 PM", "11:00 PM"},
	Late:       {"11:00 PM", "2:00 AM"},
	Overnight:  {"2:00 AM", "4:00 AM"},
	All:        {"12:00 AM", "12:00 PM"},
}

type Scheduler struct {
	Content struct {
		Programs []*Program
	}
}

func (s *Scheduler) Run() error {
	log.Info("Starting Daemon")
	now := time.Now()
	ts := getTimeSlot(&now)
	// if randomized mode do x

	for _, p := range s.Content.Programs {
		log.Debugf("program %v", formatter.StructToIndentedString(p))
		// Check Timeslots
		if ts == p.Timeslot || ts == All {
			log.Infof("getting media type: %v", p.Type)
			content := p.GetMedia()
			log.Debugf("media struct: %v", content)
			content.Get()
			// TODO go routine to check for interrupts
			content.Play() // play will block until done
		}
	}
	// if radio station start 30 minute counter.
	// smartly allocate programs to timeslots based on length if known
	// if time between TimeSlotMap
	// play program from that slot.
	// wait for program to finish
	// get next
	return nil
}

// Shuffle plays through the config content at random
func (s *Scheduler) Shuffle() error {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s.Content.Programs),
		func(i, j int) {
			s.Content.Programs[i], s.Content.Programs[j] = s.Content.Programs[j], s.Content.Programs[i]
		})
	for _, p := range s.Content.Programs {
		log.Debugf("program %v", formatter.StructToIndentedString(p))
		log.Infof("getting media type: %v", p.Type)
		content := p.GetMedia()
		log.Debugf("media struct: %v", content)
		content.Get()
		// TODO go routine to check for interrupts
		content.Play() // play will block until done
	}
	return nil
}

func getTimeSlot(t *time.Time) Timeslot {
	// if t between certain times return Timeslot
	return All
}

func NewScheduler(file string) *Scheduler {
	log.Info("Loading Config File from: ", file)
	viper.SetConfigType("yaml")
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		log.WithField("file", file).WithError(err).Error("Failed to unmarshal config file")
		return nil
	}
	s := new(Scheduler)

	if err := viper.Unmarshal(s); err != nil {
		log.WithError(err).Error("unable to unmarshall config into struct")
		return nil
	}
	log.Info("config loaded", formatter.StructToIndentedString(s))
	return s
}
