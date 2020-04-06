package store

import "github.com/cmatrixprobe/proxygool/model"

type Store interface {
	Set(address *model.Address) error
	Get(key string) (*model.Address, error)
	GetRandOne() (*model.Address, error)
	GetRandHttps() (*model.Address, error)
	GetAll() ([]*model.Address, error)
	Delete(key string) error
	Update(address *model.Address) error
	Count() (int64, error)
	Close() error
}
