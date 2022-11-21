package content

import (
	"errors"
	"fmt"
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

func (f *Folder) Get() (err error) {
	log.Infof("buffering files from %s", f.Path)
	f.Content, err = os.ReadDir(f.Path)
	if err != nil {
		return errors.New(fmt.Sprintf("unable to read folder from path: %v", err))
	}
	return nil
}

func (f *Folder) Play() error {
	// loop through the folder and play each as a local file
	for _, file := range f.Content {
		localFile, err := f.getLocalFile(file)
		if err != nil {
			return err
		}
		err = localFile.Play()
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Folder) Stop() error {
	log.Infof("Stopping stream from %v ", f.Path)
	return nil
}

func (f *Folder) getLocalFile(file os.DirEntry) (*LocalFile, error) {
	localFile := &LocalFile{
		Name: file.Name(),
		Path: f.Path + "/" + file.Name(),
	}
	err := localFile.Get()
	if err != nil {
		return nil, err
	}
	return localFile, nil
}
