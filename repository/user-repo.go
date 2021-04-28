package repository

import (
	"golang-echo-redis-cache-rest-api-example/model"
	"golang-echo-redis-cache-rest-api-example/util"
)

type UserRepository interface {
	Count() int64
	FindAll() ([]*model.User, error)
	FindAllWithPagination(page int64, limit int64) (*util.PagedModel, error)
	Save(user *model.User) (*model.User, error)
	FindById(id string) (*model.User, error)
	Update(id string, user *model.User) (*model.User, error)
	DeleteById(id string) error
}
