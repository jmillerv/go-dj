package content

//go:generate mockgen -source ./$GOFILE -destination ./mocks/mock_$GOFILE -package $GOPACKAGE

// content type should be able to be set from the configuration

const (
	podcastContent      MediaType = "podcast"
	announcementContent MediaType = "announcement"
	webRadioContent     MediaType = "web_radio"
	fileContent         MediaType = "file"
	folderContent       MediaType = "folder"
)

type MediaType string

var MediaTypeMap = map[MediaType]Media{
	podcastContent:      new(Podcast),
	announcementContent: new(Announcement),
	webRadioContent:     new(WebRadio),
	fileContent:         new(LocalFile),
	folderContent:       new(Folder),
}

type Media interface {
	Get() error
	Play() error
	Stop()
}
