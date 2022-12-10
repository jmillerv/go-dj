package content

// This package uses a different library under the hood than beep for playing radio content.
// The plan is to eventually make a custom beep.StreamCloser for the use case of an infinite radio stream
// inspiration for this solution came from https://github.com/jcheng8/goradio

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

const webRadioPlayerName = "mpv"

type WebRadio struct {
	Name   string
	URL    string
	Player streamPlayer
}

var webRadioStream streamPlayer

func (w *WebRadio) Get() error {
	var err error

	// setup web radio stream
	webRadioStream.playerName = webRadioPlayerName
	webRadioStream.url = w.URL
	webRadioStream.command = exec.Command(webRadioStream.playerName, "-quiet", "-playlist", webRadioStream.url)
	webRadioStream.in, err = webRadioStream.command.StdinPipe()
	if err != nil {
		return errors.Wrap(err, "error creating standard pipe in")
	}
	webRadioStream.out, err = webRadioStream.command.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, "error creating standard pipe out")
	}
	webRadioStream.isPlaying = false
	w.Player = webRadioStream
	return nil
}

func (w *WebRadio) Play() error {
	log.Infof("streaming from %v ", w.URL)
	if !w.Player.isPlaying {
		err := w.Player.command.Start()
		if err != nil {
			return errors.Wrap(err, "error starting web radio webRadioPlayerName")
		}
		w.Player.isPlaying = true
		done := make(chan bool)
		func() {
			w.Player.pipeChan <- w.Player.out
			done <- true
		}()
		<-done
	}
	return nil
}

func (w *WebRadio) Stop() error {
	log.Infof("Stopping stream from %v ", w.URL)
	if w.Player.isPlaying {
		w.Player.isPlaying = false
		_, err := w.Player.in.Write([]byte("q"))
		if err != nil {
			log.WithError(err).Error("error stopping web radio webRadioPlayerName: w.Player.in.Write()")
		}
		err = w.Player.in.Close()
		if err != nil {
			log.WithError(err).Error("error stopping web radio webRadioPlayerName: w.Player.in.Close()")
		}
		err = w.Player.out.Close()
		if err != nil {
			log.WithError(err).Error("error stopping web radio webRadioPlayerName: w.Player.out.Close()")
		}
		w.Player.command = nil

		w.Player.url = ""
	}
	return nil
}
