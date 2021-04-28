package service

import (
	uuid "github.com/satori/go.uuid"
	"golang-echo-redis-cache-rest-api-example/model"
	"golang-echo-redis-cache-rest-api-example/repository"
	"golang-echo-redis-cache-rest-api-example/util"
)

type UserService interface {
	FindAll(page int64, limit int64) (*util.PagedModel, error)
	Save(user *model.User) (*model.User, error)
	FindById(id string) (*model.User, error)
	Update(id string, user *model.User) (*model.User, error)
	DeleteById(id string) error
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{userRepo: userRepo}
}

func (userService *userServiceImpl) FindAll(page int64, limit int64) (*util.PagedModel, error) {
	return userService.userRepo.FindAllWithPagination(page, limit)
}

func (userService *userServiceImpl) Save(user *model.User) (*model.User, error) {
	id := uuid.NewV4().String()
	user.ID = id
	return userService.userRepo.Save(user)
}

func (userService *userServiceImpl) FindById(id string) (*model.User, error) {
	return userService.userRepo.FindById(id)
}

func (userService *userServiceImpl) Update(id string, user *model.User) (*model.User, error) {
	return userService.userRepo.Update(id, user)
}

func (userService *userServiceImpl) DeleteById(id string) error {
	return userService.userRepo.DeleteById(id)
}
