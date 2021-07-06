package content

// Not yet implemented

type LocalFile struct {
	Name    string
	Content []byte
	Path    string
}

func (l *LocalFile) Get() {
	panic("implement me")
}

func (l *LocalFile) Play() {
	panic("implement me")
}

func (l *LocalFile) Stop() {
	panic("implement me")
}
