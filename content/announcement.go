package content

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Not yet implemented

type Announcement struct {
	Name    string
	Content []byte
	URL     string
	Path    string
} // it may be possible to simply make announcement use the file struct and add a URL to the file struct

func (a *Announcement) Get() error {
	return errors.Wrap(errors.New("test"), "implement me")
}

func (a *Announcement) Play() error {
	return errors.Wrap(errors.New("test"), "implement me")
}

func (a *Announcement) Stop() error {
	log.Infof("Stopping stream from %v ", a.Path)
	return nil
}
