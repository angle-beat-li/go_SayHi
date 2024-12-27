package cache

import "github.com/goburrow/cache"

type userCache struct {
	cache            cache.LoadingCache
	scoreRankCache   cache.LoadingCache
	checkInRankCache cache.LoadingCache
}
