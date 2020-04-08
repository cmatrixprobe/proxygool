package cache

import (
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/cmatrixprobe/proxygool/util"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

// RedisStore is an default store for proxy address infos.
type RedisStore struct {
	proxyPool string
	proxyInfo string
	pool      *redis.Pool
}

// NewRedisStore returns a new standard redis store for proxies.
func NewRedisStore() *RedisStore {
	redisConfig := viper.Sub("redis")
	proxyPool := redisConfig.GetString("proxyPool")
	proxyInfo := redisConfig.GetString("proxyInfo")
	return &RedisStore{
		proxyPool: proxyPool,
		proxyInfo: proxyInfo,
		pool:      GetPool(),
	}
}

// Set .
func (s *RedisStore) Set(address *model.Address) error {
	c := s.pool.Get()
	defer c.Close()

	addr := util.CombAddr(address)
	_, err := c.Do("SADD", s.proxyPool, addr)
	if err != nil {
		return err
	}
	_, err = c.Do("HSET", s.proxyInfo, addr, util.AddressMarshal(address))
	return err
}

// Get .
func (s *RedisStore) Get(key string) (*model.Address, error) {
	c := s.pool.Get()
	defer c.Close()

	res, err := redis.String(c.Do("HGET", s.proxyInfo, key))
	if err != nil {
		return nil, err
	}
	return util.AddressUnMarshal(res), nil
}

// GetRandOne .
func (s *RedisStore) GetRandOne() (*model.Address, error) {
	c := s.pool.Get()
	defer c.Close()

	key, err := redis.String(c.Do("SRANDMEMBER", s.proxyPool))
	if err != nil {
		return nil, err
	}
	res, err := redis.String(c.Do("HGET", s.proxyInfo, key))
	if err != nil {
		return nil, err
	}
	return util.AddressUnMarshal(res), nil
}

// GetRandHTTPS .
func (s *RedisStore) GetRandHTTPS() (*model.Address, error) {
	c := s.pool.Get()
	defer c.Close()

	addresses, err := s.GetAll()
	if err != nil {
		return nil, err
	}
	var httpsAddr []*model.Address
	for _, addr := range addresses {
		if addr.Protocol == "https" {
			httpsAddr = append(httpsAddr, addr)
		}
	}
	return util.RandomElement(httpsAddr), nil
}

// GetAll .
func (s *RedisStore) GetAll() ([]*model.Address, error) {
	c := s.pool.Get()
	defer c.Close()

	var addresses []*model.Address
	reply, err := redis.Strings(c.Do("HGETALL", s.proxyInfo))
	if err != nil {
		return nil, err
	}
	for i := range reply {
		if i&1 == 1 {
			addresses = append(addresses, util.AddressUnMarshal(reply[i]))
		}
	}
	return addresses, nil
}

// Delete .
func (s *RedisStore) Delete(key string) error {
	c := s.pool.Get()
	defer c.Close()

	_, err := c.Do("SREM", s.proxyPool, key)
	if err != nil {
		return err
	}
	_, err = c.Do("HDEL", s.proxyInfo, key)
	return err
}

// Update .
func (s *RedisStore) Update(address *model.Address) error {
	c := s.pool.Get()
	defer c.Close()

	addr := util.CombAddr(address)
	_, err := c.Do("HSET", s.proxyInfo, addr, util.AddressMarshal(address))
	return err
}

// Count .
func (s *RedisStore) Count() (int64, error) {
	c := s.pool.Get()
	defer c.Close()

	return redis.Int64(c.Do("SCARD", s.proxyPool))
}

// Close .
func (s *RedisStore) Close() error {
	return s.pool.Close()
}
