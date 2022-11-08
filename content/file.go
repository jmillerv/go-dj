package content

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

type LocalFile struct {
	Name    string
	Content *os.File
	Path    string
}

func (l *LocalFile) Get() error {
	log.Infof("buffering file from %s", l.Path)
	f, err := os.Open(l.Path)
	if err != nil {
		return errors.Wrap(err, "unable to open file from path")
	}
	l.Content = f
	return nil
}

func (l *LocalFile) Play() error {
	log.Infof("decoding file from %v", l.Path)
	streamer, format, err := mp3.Decode(l.Content)
	if err != nil {
		return errors.Wrap(err, "mp3.Decode: unable to decode mp3")
	}
	defer streamer.Close()
	log.Infof("playing file buffer from %v", l.Path)
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return errors.Wrap(err, "speaker.Init: unable to play file")
	}
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
	return nil
}

func (l *LocalFile) Stop() {
	log.Infof("Stopping stream from %v ", l.Path)
}
