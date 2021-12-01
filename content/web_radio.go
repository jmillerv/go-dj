package content

// This package uses a different library under the hood than beep for playing radio content.
// The plan is to eventually make a custom beep.StreamCloser for the use case of an infinite radio stream
// inspiration for this solution came from https://github.com/jcheng8/goradio

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os/exec"
)

const player = "mpv"

type WebRadio struct {
	Name   string
	URL    string
	Player wrCommand
}

type wrCommand struct {
	playerName string
	url        string
	isPlaying  bool
	command    *exec.Cmd
	in         io.WriteCloser
	out        io.ReadCloser
	pipeChan   chan io.ReadCloser
}

var wrc wrCommand

func (w *WebRadio) Get() {
	var err error
	wrc.playerName = player
	wrc.url = w.URL
	wrc.command = exec.Command(wrc.playerName, "-quiet", "-playlist", wrc.url)
	wrc.in, err = wrc.command.StdinPipe()
	if err != nil {
		log.WithError(err).Error("error creating standard pipe in")
	}
	wrc.out, err = wrc.command.StdoutPipe()
	if err != nil {
		log.WithError(err).Error("error creating standard pipe out")
	}
	wrc.isPlaying = false
	w.Player = wrc
}

func (w *WebRadio) Play() {
	log.Infof("streaming from URL %v ", w.URL)
	if !w.Player.isPlaying {
		err := w.Player.command.Start()
		if err != nil {
			log.WithError(err).Error("error starting web radio player")
		}
		w.Player.isPlaying = true
		done := make(chan bool)
		func() {
			w.Player.pipeChan <- w.Player.out
			done <- true
		}()
		<-done
	}

}

func (w *WebRadio) Stop() {
	panic("implement me")
}
