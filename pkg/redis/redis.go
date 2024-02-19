package redis

import (
	"context"
	"crypto/tls"
	"errors"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

const redisTimeoutPing = 5 * time.Second

var ErrRedisConnection = errors.New("failed connection to redis")

type RInstance struct {
	Client *redis.Client
}

var (
	redisClientOnce sync.Once
	redisInstance   *RInstance
)

func Instance(host, password string, db int, tlsConf bool) (*RInstance, error) {
	if redisInstance != nil {
		return redisInstance, Ping(redisInstance.Client)
	}
	redisClientOnce.Do(func() { redisInstance = &RInstance{Client: InstanceConnect(host, password, db, tlsConf)} })
	return redisInstance, Ping(redisInstance.Client)
}

func InstanceConnect(host, password string, db int, tlsConf bool) *redis.Client {
	connOptions := &redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	}
	if tlsConf {
		connOptions.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
	}
	return redis.NewClient(connOptions)
}

func Ping(rConn *redis.Client) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), redisTimeoutPing)
	defer cancelFunc()
	if err := rConn.Ping(ctx).Err(); err != nil {
		return ErrRedisConnection
	}
	return nil
}
