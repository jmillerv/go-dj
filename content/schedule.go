package content

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

type Schedule struct {
	Early      bool
	Morning    bool
	Breakfast  bool
	Midmorning bool
	Afternoon  bool
	Commute    bool
	Evening    bool
	Late       bool
	Overnight  bool
}

type Scheduler struct {
	Programs []*Program
	Schedule Schedule
}
