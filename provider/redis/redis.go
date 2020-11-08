package redis

import (
	"jettster/provider/config"

	"github.com/go-redis/redis"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     config.GetString("redis.url"),
	Password: config.GetString("redis.password"),
	DB:       config.GetInt("redis.db"),
})

func Set(key string, value string) {
	err := rdb.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func Get(key string) (string, bool) {
	val, err := rdb.Get(key).Result()
	if err != nil {
		return "", false
	}
	return val, true
}
