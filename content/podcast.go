package content

type Podcast struct {
	Name string
	URL  string
	File []byte
}

func (p Podcast) Get() {
	panic("implement me")
}

func (p Podcast) Play() {
	panic("implement me")
}

func (p Podcast) Stop() {
	panic("implement me")
}
