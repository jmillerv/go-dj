package content

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

type LocalFile struct {
	Name     string
	Content  *os.File
	Path     string
	streamer beep.StreamSeekCloser
}

func (l *LocalFile) Get() {
	log.Infof("buffering file from %s", l.Path)
	f, err := os.Open(l.Path)
	if err != nil {
		log.WithError(err).Error("unable to open file from path")
	}
	l.Content = f
}

func (l *LocalFile) Play() {
	log.Infof("decoding file from %v", l.Path)
	streamer, format, err := mp3.Decode(l.Content)
	if err != nil {
		log.WithError(err).Fatal("unable to decode mp3")
	}
	l.streamer = streamer
	defer l.streamer.Close()
	log.Infof("playing file buffer from %v", l.Path)
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.WithError(err).Fatal("unable to play file")
	}
	done := make(chan bool)
	speaker.Play(beep.Seq(l.streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}

func (l *LocalFile) Stop() {
	log.Infof("Stopping stream from %v ", l.Path)
}
