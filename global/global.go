package global

import (
	"github.com/cmatrixprobe/proxygool/cache"
	"github.com/cmatrixprobe/proxygool/store"
)

var globalStore store.Store

func init() {
	// default storage: redis
	globalStore = cache.NewRedisStore()
}

func GetStore() store.Store {
	return globalStore
}

func SetCustomStore(s store.Store) {
	globalStore = s
}
