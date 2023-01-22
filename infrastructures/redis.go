package infrastructures

import (
	"github.com/corentings/kafejo/config"
	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
)

// ConnectRedis Connects to redis
func ConnectRedis() error {
	redisClient = NewRedisClient()

	_, err := redisClient.Ping(redisClient.Context()).Result()
	if err != nil {
		return err
	}

	return nil
}

// CloseRedis Closes redis connection
func CloseRedis() error {
	return redisClient.Close()
}

// GetRedisClient Returns redis client
func GetRedisClient() *redis.Client {
	return redisClient
}

// NewRedisClient Returns new redis client
func NewRedisClient() *redis.Client {
	redisHost := config.RedisHost

	client := redis.NewClient(&redis.Options{
		Addr:         redisHost,
		MinIdleConns: config.RedisMinIdleConns,
		PoolSize:     config.RedisPoolSize,
		PoolTimeout:  config.RedisPoolTimeout,
	})

	return client
}
