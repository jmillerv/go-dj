package content

import log "github.com/sirupsen/logrus"

type Podcast struct {
	Name    string
	URL     string
	Path    string
	Content []byte
}

func (p *Podcast) Get() error {
	panic("implement me")
}

func (p *Podcast) Play() error {
	panic("implement me")
}

func (p *Podcast) Stop() {
	log.Infof("Stopping stream from %v ", p.Path)
}
