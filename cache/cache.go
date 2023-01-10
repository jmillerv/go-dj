package cache

import (
	"os"

	"zgo.at/zcache"
)

const (
	defaultPodcastCache   = "podcastCache"
	podcastCacheLocalFile = "./cache/podcastCache.json"
)

var PodcastPlayedCache *zcache.Cache

func ClearPodcastPlayedCache() error {
	PodcastPlayedCache.Flush()
	err := os.Remove(podcastCacheLocalFile)
	if err != nil {
		return err
	}
	return nil
}
