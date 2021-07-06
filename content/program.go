package content

type Program struct {
	Name     string
	Source   string
	Timeslot Slot
	Type     MediaType
}

func (p *Program) Get() *Media {
	media := p.mediaFactory()
	return &media
}

func (p *Program) mediaFactory() Media {
	m := MediaTypeMap[p.Type]
	switch m.(type) {
	case *Announcement:
		panic("implement me")
	case *LocalFile:
		panic("implement me")
	case *Podcast:
		pod := m.(*Podcast)
		pod.Name = p.Name
		pod.URL = p.Source
		return pod
	case *WebRadio:
		panic("implement me")
	}
	return nil
}
