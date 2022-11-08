package content

// This package uses a different library under the hood than beep for playing radio content.
// The plan is to eventually make a custom beep.StreamCloser for the use case of an infinite radio stream
// inspiration for this solution came from https://github.com/jcheng8/goradio

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"os/exec"
)

const player = "mpv"

type WebRadio struct {
	Name   string
	URL    string
	Player webRadioCommand
}

type webRadioCommand struct {
	playerName string
	url        string
	isPlaying  bool
	command    *exec.Cmd
	in         io.WriteCloser
	out        io.ReadCloser
	pipeChan   chan io.ReadCloser
}

var wrc webRadioCommand

func (w *WebRadio) Get() error {
	var err error
	wrc.playerName = player
	wrc.url = w.URL
	wrc.command = exec.Command(wrc.playerName, "-quiet", "-playlist", wrc.url)
	wrc.in, err = wrc.command.StdinPipe()
	if err != nil {
		return errors.Wrap(err, "error creating standard pipe in")
	}
	wrc.out, err = wrc.command.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, "error creating standard pipe out")
	}
	wrc.isPlaying = false
	w.Player = wrc
	return nil
}

func (w *WebRadio) Play() error {
	log.Infof("streaming from %v ", w.URL)
	if !w.Player.isPlaying {
		err := w.Player.command.Start()
		if err != nil {
			return errors.Wrap(err, "error starting web radio player")
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

func (w *WebRadio) Stop() {
	log.Infof("Stopping stream from %v ", w.URL)
	if w.Player.isPlaying {
		w.Player.isPlaying = false
		_, err := w.Player.in.Write([]byte("q"))
		if err != nil {
			log.WithError(err).Error("error stopping web radio player: w.Player.in.Write()")
		}
		err = w.Player.in.Close()
		if err != nil {
			log.WithError(err).Error("error stopping web radio player: w.Player.in.Close()")
		}
		err = w.Player.out.Close()
		if err != nil {
			log.WithError(err).Error("error stopping web radio player: w.Player.out.Close()")
		}
		w.Player.command = nil

		w.Player.url = ""
	}
}
