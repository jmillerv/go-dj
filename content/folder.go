package content

import "os"

// The folder structure exists because I didn't want to load an entire folder's worth
// of songs into memory like the LocalFile struct does.

type Folder struct {
	Name    string
	Content []*os.DirEntry
	Path    string
}

func (f *Folder) Get() {
	// load the files into the Content
	panic("implement me")
}

func (f *Folder) Play() {
	// loop through the folder and play each as a local file
	panic("implement me")
}

func (f *Folder) Stop() {
	panic("implement me")
}

func (f *Folder) getLocalFile() *LocalFile {
	return nil
}
