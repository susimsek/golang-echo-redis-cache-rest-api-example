package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"golang-echo-redis-cache-rest-api-example/config"
	"golang-echo-redis-cache-rest-api-example/model"
	"time"
)

var ctx context.Context = context.Background()

type userCacheImpl struct {
	client  *redis.Client
	expires time.Duration
}

func (userCache *userCacheImpl) Delete(key string) error {
	return userCache.client.Del(ctx, key).Err()
}

func NewUserRedisCache(client *redis.Client) UserCache {
	expires, _ := time.ParseDuration(config.UserCacheExpirationMs)
	return &userCacheImpl{
		client:  client,
		expires: expires,
	}
}

func (userCache *userCacheImpl) Set(key string, value *model.User) {
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	userCache.client.Set(ctx, key, json, userCache.expires*time.Millisecond)
}

func (userCache *userCacheImpl) Get(key string) *model.User {

	val, err := userCache.client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	user := model.User{}
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		panic(err)
	}
	return &user
}
