package content

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
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
	wavFile         string = "wav"
	mp3File         string = "mp3"
	oggFile         string = "oggs"
	flacFile        string = "flac"
	sampleRateTime         = 10
	wavFileBitRate  int64  = 1441
	mp3FileBitRate  int64  = 128
	oggFileBitRate  int64  = 192
	flacFileBitRate int64  = 320
	durationBase           = 10
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
		return errors.New(fmt.Sprintf("unable to open file from path: %v", err)) //nolint:lll,revive,gosimple,nolintlint // error pref
	}

	log.Infof("decoding file from %v", l.Path)
	l.Content = f

	return nil
}

func (l *LocalFile) Play() error {
	var streamer beep.StreamSeekCloser
	var format beep.Format //nolint:wsl // it's fine for declarations to touch

	err := l.setDecoder()
	if err != nil {
		return errors.New(fmt.Sprintf("error setting decoder: %v", err)) //nolint:revive,gosimple // error pref
	}

	_, err = l.Content.Seek(0, 0)
	if err != nil {
		return errors.New(fmt.Sprintf("unable to seek to beginning of file: %v", err)) //nolint:lll,revive,gosimple,nolintlint // error pref
	}

	if l.fileType == wavFile || l.fileType == flacFile {
		streamer, format, err = l.decodeReader(l.Content)
		if err != nil {
			return errors.New(fmt.Sprintf("unable to decode file: %v", err)) //nolint:revive,gosimple // error pref
		}
	}

	if l.fileType == mp3File || l.fileType == oggFile {
		streamer, format, err = l.decodeReadCloser(l.Content)
		if err != nil {
			log.WithError(err).Fatal("unable to decode file")
		}
	}

	log.Infof("playing file buffer from %v", l.Path)

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/sampleRateTime))
	if err != nil {
		return errors.New(fmt.Sprintf("unable to play file: %v", err)) //nolint:revive,gosimple,nolintlint // error pref
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() { //nolint:wsl // grouping makes sense here
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
		return errors.New("empty bytes") //nolint:goerr113 // we want this error as it is.
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

		return errors.New("unsupported filetype " + unknownType.Extension) //nolint:goerr113 // desired error
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

// getEstimatedFileDuration uses oft use bitrates for given filtypes and provides and estimated duration in seconds
func (l *LocalFile) getEstimatedFileDuration() string {
	// assume a constant bitrate for the file types
	switch l.fileType {
	case wavFile:
		stats, err := l.Content.Stat()
		if err != nil {
			log.WithError(err).Error("failed to get wav file stats")
			return ""
		}
		duration := stats.Size() / wavFileBitRate

		return strconv.FormatInt(duration, durationBase)

	case mp3File:
		stats, err := l.Content.Stat()
		if err != nil {
			log.WithError(err).Error("failed to get mp3 file stats")
			return ""
		}
		duration := stats.Size() / mp3FileBitRate

		return strconv.FormatInt(duration, durationBase)

	case oggFile:
		stats, err := l.Content.Stat()
		if err != nil {
			log.WithError(err).Error("failed to get mp3 file stats")
			return ""
		}
		duration := stats.Size() / oggFileBitRate

		return strconv.FormatInt(duration, durationBase)

	case flacFile:
		stats, err := l.Content.Stat()
		if err != nil {
			log.WithError(err).Error("failed to get FLAC file stats")
			return ""
		}
		duration := stats.Size() / flacFileBitRate

		return strconv.FormatInt(duration, durationBase)

	default:
		return "unknown file type: can't determine duration"
	}
}
