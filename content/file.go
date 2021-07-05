package content

type LocalFile struct {
	Content []byte
	Path    string
}

func (f *LocalFile) Get() {
	panic("implement me")
}

func (f *LocalFile) Play() {
	panic("implement me")
}

func (f *LocalFile) Stop() {
	panic("implement me")
}
