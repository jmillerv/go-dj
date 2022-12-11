package content

import (
	"io"
	"os/exec"
)

type streamPlayer struct {
	playerName string
	url        string
	isPlaying  bool
	command    *exec.Cmd
	in         io.WriteCloser
	out        io.ReadCloser
	pipeChan   chan io.ReadCloser
}
