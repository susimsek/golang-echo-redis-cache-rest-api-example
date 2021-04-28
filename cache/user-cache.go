package cache

import (
	"golang-echo-redis-cache-rest-api-example/model"
)

type UserCache interface {
	Set(key string, value *model.User)
	Get(key string) *model.User
	Delete(key string) error
}
