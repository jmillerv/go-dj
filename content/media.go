package content

type Media interface {
	Get()
	Play()
	Stop()
}

// Media factory takes in a an array of sources and returns the media objects.
func MediaFactory() []*Media {
	return nil
}
