package configs

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Db struct{}

var Client *redis.Client
var Ctx = context.Background()

func InitRedis(config Config) {

	fmt.Println("Redis starting...")
	dsn := config.Redis.Dsn

	if len(dsn) == 0 {
		dsn = "redis:6379"
	}

	Client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err := Client.Ping(Ctx).Result()

	if err != nil {
		panic(err)
	}
}

func GetRedis() *redis.Client {
	return Client
}
