package content

const (
	podcastContent      MediaType = "podcast"
	announcementContent MediaType = "announcement"
	webRadioContent     MediaType = "web_radio"
	fileContent         MediaType = "file"
)

type MediaType string

var MediaTypeMap = map[MediaType]Media{
	podcastContent:      new(Podcast),
	announcementContent: new(Announcement),
	webRadioContent:     new(WebRadio),
	fileContent:         new(LocalFile),
}

type Media interface {
	Get()
	Play()
	Stop()
}
