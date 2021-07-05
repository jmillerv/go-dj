package content

type Announcement struct {
	Content []byte
	Path    string
}

func (a *Announcement) Get() {
	panic("implement me")
}

func (a *Announcement) Play() {
	panic("implement me")
}

func (a *Announcement) Stop() {
	panic("implement me")
}
