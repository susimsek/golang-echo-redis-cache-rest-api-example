package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

func RedisConnection() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     RedisUrl,
		Password: RedisPassword,
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalf("Cannot connect to Redis : %s\n", err)
	} else {
		fmt.Printf("We are connected to the Redis\n")
	}

	return client
}
