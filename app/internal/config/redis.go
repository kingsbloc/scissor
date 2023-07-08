package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/kingsbloc/scissor/internal/utils"
)

var RDB *redis.Client

func ConnectRedis() *redis.Client {
	uri := utils.GetEnvVar("REDIS_URI")
	opt, err := redis.ParseURL(uri)
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)
	_, err = rdb.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
	log.Println("==== Redis Connected ====")
	RDB = rdb
	return rdb
}
