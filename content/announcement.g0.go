package content

type Announcement struct {
	File []byte
}

func (a Announcement) Get() {
	panic("implement me")
}

func (a Announcement) Play() {
	panic("implement me")
}

func (a Announcement) Stop() {
	panic("implement me")
}
