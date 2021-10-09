package content

// Not yet implemented

type Announcement struct {
	Name    string
	Content []byte
	URL     string
	Path    string
} // it may be possible to simply make announcement use the file struct and add a URL to the file struct

func (a *Announcement) Get() {
	panic("implement me")
}

func (a *Announcement) Play() {
	panic("implement me")
}

func (a *Announcement) Stop() {
	panic("implement me")
}
