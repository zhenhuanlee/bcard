package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

// RedisConn redis pool
var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()

	if err != nil {
		panic("init redis connection failed")
	}

	fmt.Println(pong)
}

// Get get
func Get(key string) (string, error) {
	return client.Get(key).Result()
}
