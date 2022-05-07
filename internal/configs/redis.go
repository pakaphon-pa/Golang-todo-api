package configs

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type Db struct{}

var Client *redis.Client
var ctx = context.Background()

func InitRedis() {
	fmt.Println("Redis starting...")
	dsn := viper.GetString(`redis.dsn`)

	if len(dsn) == 0 {
		dsn = "redis:6379"
	}

	Client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err := Client.Ping(ctx).Result()

	if err != nil {
		panic(err)
	}
}

func GetRedis() *redis.Client {
	return Client
}
