package cache

import (
	_ "github.com/cmatrixprobe/proxygool/config"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"time"
)

var pool *redis.Pool

func init() {
	pool = newPool()
}

func newPool() *redis.Pool {
	// Read configuration from application.yaml
	redisConfig := viper.Sub("redis")
	network := redisConfig.GetString("network")
	password := redisConfig.GetString("password")
	MaxIdle := redisConfig.GetInt("MaxIdle")
	MaxActive := redisConfig.GetInt("MaxActive")
	IdleTimeout := redisConfig.GetDuration("IdleTimeout")
	testFrequency := redisConfig.GetDuration("testFrequency")
	Wait := redisConfig.GetBool("Wait")
	ip := redisConfig.GetString("ip")
	port := redisConfig.GetString("port")
	address := ip + ":" + port

	return &redis.Pool{
		MaxIdle:     MaxIdle,
		MaxActive:   MaxActive,
		IdleTimeout: IdleTimeout,
		Wait:        Wait,
		Dial: func() (c redis.Conn, err error) {
			c, err = redis.Dial(network, address, redis.DialPassword(password))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < testFrequency {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func Pool() *redis.Pool {
	return pool
}
