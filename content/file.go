package content

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
	"github.com/h2non/filetype"
	log "github.com/sirupsen/logrus"
)

const (
	wavFile  = "wav"
	mp3File  = "mp3"
	oggFile  = "oggs"
	flacFile = "flac"
)

type LocalFile struct {
	Name             string
	Content          *os.File
	Path             string
	decodeReader     func(r io.Reader) (s beep.StreamSeekCloser, format beep.Format, err error)
	decodeReadCloser func(rc io.ReadCloser) (s beep.StreamSeekCloser, format beep.Format, err error)
	fileType         string
}

func (l *LocalFile) Get() error {
	log.Infof("buffering file from %s", l.Path)
	f, err := os.Open(l.Path)
	if err != nil {
		return errors.New(fmt.Sprintf("unable to open file from path: %v", err))
	}
	log.Infof("decoding file from %v", l.Path)
	l.Content = f
	return nil
}

func (l *LocalFile) Play() error {
	var streamer beep.StreamSeekCloser
	var format beep.Format
	err := l.setDecoder()
	if err != nil {
		return errors.New(fmt.Sprintf("error setting decoder: %v", err))
	}
	_, err = l.Content.Seek(0, 0)
	if err != nil {
		return errors.New(fmt.Sprintf("unable to seek to beginning of file: %v", err))
	}
	if l.fileType == wavFile || l.fileType == flacFile {
		streamer, format, err = l.decodeReader(l.Content)
		if err != nil {
			return errors.New(fmt.Sprintf("unable to decode file: %v", err))
		}
	}
	if l.fileType == mp3File || l.fileType == oggFile {
		streamer, format, err = l.decodeReadCloser(l.Content)
		if err != nil {
			log.WithError(err).Fatal("unable to decode file")
		}
	}
	log.Infof("playing file buffer from %v", l.Path)
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return errors.New(fmt.Sprintf("unable to play file: %v", err))
	}
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
	return nil
}

func (l *LocalFile) Stop() error {
	log.Infof("file.Stop::Stopping stream from %v ", l.Path)
	return nil
}

func (l *LocalFile) setDecoder() error {
	buf, err := io.ReadAll(l.Content)
	if buf == nil {
		return errors.New("empty bytes")
	}
	if err != nil {
		return err
	}
	switch l.getFileType(buf) {
	case wavFile:
		l.fileType = wavFile
		l.decodeReader = wav.Decode
	case flacFile:
		l.fileType = flacFile
		l.decodeReader = flac.Decode
	case mp3File:
		l.fileType = mp3File
		l.decodeReadCloser = mp3.Decode
	case oggFile:
		l.fileType = oggFile
		l.decodeReadCloser = vorbis.Decode
	default:
		err := l.Stop()
		if err != nil {
			log.WithError(err).Error("localFile.setDecoder::error stopping local file")
		}
		unknownType, err := filetype.Match(buf)
		if err != nil {
			log.WithError(err).Error("localFile.setDecoder::error getting filetype")
		}
		return errors.New("unsupported filetype " + unknownType.Extension)
	}
	return nil
}
func (l *LocalFile) getFileType(buf []byte) string {
	ext := filepath.Ext(l.Path)
	trimmedExt := strings.TrimLeft(ext, ".") // remove the delimiter
	// added the trim check because some supported filetypes were not recognized by
	// the filetype.IsType function despite having proper extension and working with the respective decoder
	if filetype.IsType(buf, filetype.GetType("wav")) || trimmedExt == wavFile {
		return wavFile
	}
	if filetype.IsType(buf, filetype.GetType("mp3")) || trimmedExt == mp3File {
		return mp3File
	}
	if filetype.IsType(buf, filetype.GetType("ogg")) || trimmedExt == oggFile {
		return oggFile
	}
	if filetype.IsType(buf, filetype.GetType("flac")) || trimmedExt == flacFile {
		return flacFile
	}
	return ""
}
