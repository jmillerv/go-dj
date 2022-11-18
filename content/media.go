package content

import (
	"github.com/faiface/beep"
	"io"
)

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

// Media is the interface to represent playing any type of audio.
type Media interface {
	Get()
	Play()
	Stop()
}

// Decoder is an interface to the beep package.
type Decoder interface {
	Decode(rc io.ReadCloser) (s beep.StreamSeekCloser, format beep.Format, err error)
}
