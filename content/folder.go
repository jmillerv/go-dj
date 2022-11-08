package content

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
)

// The folder structure exists because I didn't want to load an entire folder's worth
// of songs into memory like the LocalFile struct does.

// Folder is a struct for parsing folders that implements the Media interface
type Folder struct {
	Name    string
	Content []os.DirEntry
	Path    string
}

func (f *Folder) Get() error {
	log.Infof("buffering files from %s", f.Path)
	folder, err := os.ReadDir(f.Path)
	if err != nil {
		return errors.Wrap(err, "unable to read folder from path")
	}
	f.Content = folder
	return nil
}

func (f *Folder) Play() error {
	// loop through the folder and play each as a local file
	for _, file := range f.Content {
		l := f.getLocalFile(file)
		err := l.Play()
		if err != nil {
			return errors.Wrap(err, "unable to play file from folder")
		}
	}
	return nil
}

func (f *Folder) Stop() {
	log.Infof("Stopping stream from %v ", f.Path)
}

func (f *Folder) getLocalFile(file os.DirEntry) *LocalFile {
	l := &LocalFile{
		Name: file.Name(),
		Path: f.Path + "/" + file.Name(),
	}
	l.Get()
	return l
}
