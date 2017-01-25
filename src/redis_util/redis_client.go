package redis_util

import (
	// "fmt"
	"github.com/pquerna/ffjson/ffjson"
	"gopkg.in/redis.v5"
	"log"
	"sync"
	"time"
)

type RedisClient struct {
	pool *redis.Client
}

var once sync.Once
var redisClient *RedisClient

func GetRedisClientInstance() *RedisClient {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:         "127.0.0.1:6379",
			Password:     "",
			MaxRetries:   3,
			DialTimeout:  5 * time.Second,
			WriteTimeout: 3 * time.Second,
			PoolSize:     20,
			PoolTimeout:  0,
			IdleTimeout:  0,
			DB:           3,
		})

		pong, err := client.Ping().Result()
		log.Println(pong, err)
		redisClient = &RedisClient{client}
	})
	return redisClient
}

func (c *RedisClient) Get(key string) map[string]interface{} {
	val, err := c.pool.Get(key).Result()
	log.Println(val, err)
	var dat map[string]interface{}
	if err == nil {
		ffjson.Unmarshal([]byte(val), &dat)
		if len(dat) == 0 {
			return map[string]interface{}{"val": val}
		}
	}
	return dat
}
