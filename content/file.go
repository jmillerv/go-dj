package content

import (
	"errors"
	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

const (
	wavByteHeader  = "RIFF"
	mp3ByteHeader1 = "0xFFE"
	mp3ByteHeader2 = "0XFFF"
	oggByteHeader  = "OggS"
	flacByteHeader = "fLaC"
)

type LocalFile struct {
	Name    string
	Content *os.File
	Path    string
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
	streamer, format, err := l.checkHeader()
	if err != nil {
		log.WithError(err).Fatal("unable to decode mp3")
	}
	defer streamer.Close()
	log.Infof("playing file buffer from %v", l.Path)
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.WithError(err).Fatal("unable to play file")
	}
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}

func (l *LocalFile) Stop() {
	log.Infof("Stopping stream from %v ", l.Path)
}

// Check determines the byte header and returns the proper decoder (mp3, wav, FLAC, OGG)
// other decoders can be implemented by adhering to the beep.Decode interface.
func (l *LocalFile) checkHeader() (s beep.StreamSeekCloser, format beep.Format, err error) {
	buf, err := io.ReadAll(l.Content)
	if err != nil {
		return nil, beep.Format{}, err
	}
	if err != nil {
		return nil, beep.Format{}, err
	}
	switch getByteHeader(buf) {
	case wavByteHeader:
		return wav.Decode(l.Content)
	case mp3ByteHeader1, mp3ByteHeader2:
		return mp3.Decode(l.Content)
	case oggByteHeader:
		return vorbis.Decode(l.Content)
	case flacByteHeader:
		return flac.Decode(l.Content)
	default:
		return nil, beep.Format{}, errors.New("unknown byte header")
	}
}

func getByteHeader(buf []byte) string {
	bytes := buf[4]
	return string(bytes)
}
