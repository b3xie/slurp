package internal

import "github.com/redis/go-redis/v9"

func RedisConnect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.15.178:32769",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
