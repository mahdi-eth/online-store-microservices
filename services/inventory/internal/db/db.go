package db

import (
    "github.com/go-redis/redis/v8"
    "log"
    "context"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitDB(addr string, password string, db int) {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })

    _, err := RedisClient.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Error connecting to Redis: %v", err)
    }
    log.Println("Connected to Redis")
}
