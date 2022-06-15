package content

import (
	log "github.com/sirupsen/logrus"
	"io/fs"
	"io/ioutil"
)

// The folder structure exists because I didn't want to load an entire folder's worth
// of songs into memory like the LocalFile struct does.

// Folder is a struct for parsing folders that implements the Media interface
type Folder struct {
	Name    string
	Content *[]fs.FileInfo
	Path    string
}

func (f *Folder) Get() {
	log.Infof("buffering files from %s", f.Path)
	folder, err := ioutil.ReadDir(f.Path)
	if err != nil {
		log.WithError(err).Error("unable to read folder from path")
	}
	f.Content = &folder
}

func (f *Folder) Play() {
	// loop through the folder and play each as a local file
	for _, file := range *f.Content {
		l := f.getLocalFile(file)
		l.Play()
	}
}

func (f *Folder) Stop() {
	panic("implement me")
}

func (f *Folder) getLocalFile(file fs.FileInfo) *LocalFile {
	l := &LocalFile{
		Name: file.Name(),
		Path: f.Path + "/" + file.Name(),
	}
	l.Get()
	return l
}
