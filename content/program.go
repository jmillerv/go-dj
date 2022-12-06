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

func (p *Program) GetMedia() Media {
	media := p.mediaFactory()
	return media
}

func (p *Program) mediaFactory() Media {
	m := MediaTypeMap[p.Type]
	switch m.(type) {
	case *Announcement:
		panic("implement me")
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
		panic("implement me")
	case *WebRadio:
		radio := m.(*WebRadio)
		radio.Name = p.Name
		radio.URL = p.Source
		log.Debugf("returning WebRadio: %v", formatter.StructToString(radio))
		return radio
	}
	return nil
}
