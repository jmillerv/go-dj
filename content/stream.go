package content

import (
	"io"
	"os/exec"
	"time"
)

type streamPlayer struct {
	playerName string
	url        string
	isPlaying  bool
	command    *exec.Cmd
	in         io.WriteCloser
	out        io.ReadCloser
	pipeChan   chan io.ReadCloser
	duration   time.Duration
}

func (s *streamPlayer) setDuration(duration string) {
	s.duration, _ = time.ParseDuration(duration)
}
