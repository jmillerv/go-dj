package content

import (
	"github.com/jmillerv/go-utilities/formatter"
	log "github.com/sirupsen/logrus"
)

type Program struct {
	Name     string
	Source   string
	Timeslot *Timeslot
	Type     MediaType
}

//nolint:gochecknoglobals // the globals here help but a refactor would be considered.
var (
	PodcastPlayOrderRandom   bool
	PodcastPlayerOrderOldest bool
)

func (p *Program) getMediaType() MediaType {
	return p.Type
}

func (p *Program) GetMedia() Media {
	media := p.mediaFactory()

	return media
}

//nolint:forcetypeassert,gosimple,gocritic // type is checked in the switch case
func (p *Program) mediaFactory() Media {
	m := MediaTypeMap[p.Type]
	switch m.(type) {
	case *Folder:
		folder := m.(*Folder)
		folder.Name = p.Name
		folder.Path = p.Source

		log.Debugf("returning Folder: %v", formatter.StructToString(folder))

		return folder
	case *LocalFile:
		file := m.(*LocalFile)
		file.Name = p.Name
		file.Path = p.Source

		log.Debugf("returning LocalFile: %v", formatter.StructToString(file))

		return file
	case *Podcast:
		podcast := m.(*Podcast)
		podcast.Name = p.Name
		podcast.URL = p.Source

		podcast.PlayOrder = playOrderNewest // default
		if PodcastPlayerOrderOldest == true {
			podcast.PlayOrder = playOrderOldest
		}

		if PodcastPlayOrderRandom == true {
			podcast.PlayOrder = playOrderRandom
		}

		log.Debugf("returning podcast: %v", formatter.StructToString(podcast))

		return podcast
	case *WebRadio:
		radio := m.(*WebRadio)
		radio.Name = p.Name
		radio.URL = p.Source

		log.Debugf("returning WebRadio: %v", formatter.StructToString(radio))

		return radio
	}

	return nil
}
