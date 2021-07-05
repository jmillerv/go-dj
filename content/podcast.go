package content

type Podcast struct {
	URL     string
	Path    string
	Content []byte
}

func (p *Podcast) Get() {
	panic("implement me")
}

func (p *Podcast) Play() {
	panic("implement me")
}

func (p *Podcast) Stop() {
	panic("implement me")
}
