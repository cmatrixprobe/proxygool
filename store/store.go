package store

import (
	"github.com/cmatrixprobe/proxygool/cache"
	"github.com/cmatrixprobe/proxygool/model"
)

// Store an object implementing Store interface can be
// registered with SetCustomStore to replace the default redis store.
type Store interface {
	Set(address *model.Address) error
	Get(key string) (*model.Address, error)
	GetRandOne() (*model.Address, error)
	GetRandHTTPS() (*model.Address, error)
	GetAll() ([]*model.Address, error)
	Delete(key string) error
	Update(address *model.Address) error
	Count() (int64, error)
	Close() error
}

var storage Store

func init() {
	// default storage: redis
	storage = cache.NewRedisStore()
}

// SetCustomStore replaces the default store with s.
func SetCustomStore(s Store) {
	storage = s
}
