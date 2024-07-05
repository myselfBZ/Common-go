package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitRedis() *redis.Client {
    Client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        DB: 0,
        Password: "",
    })
    _, err := Client.Ping(ctx).Result()
    if err != nil{
        panic("You suck at initializing Redis client")
    }
    return Client 
}

