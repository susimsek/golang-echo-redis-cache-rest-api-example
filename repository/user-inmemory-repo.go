package repository

import (
	"context"
	"golang-echo-redis-cache-rest-api-example/exception"
	"golang-echo-redis-cache-rest-api-example/model"
	"golang-echo-redis-cache-rest-api-example/util"
)

var ctx context.Context = context.Background()

type userRepositoryImpl struct {
	users []*model.User
}

func NewUserInMemoryRepository() UserRepository {
	return &userRepositoryImpl{
		users: make([]*model.User, 0),
	}
}

func (userRepo *userRepositoryImpl) Count() int64 {
	return int64(len(userRepo.users))
}

func (userRepo *userRepositoryImpl) FindAll() ([]*model.User, error) {
	return userRepo.users, nil
}

func (userRepo *userRepositoryImpl) FindAllWithPagination(page int64, limit int64) (*util.PagedModel, error) {
	count := int64(len(userRepo.users))
	paginator := util.Paging(page, limit, count)
	start := paginator.Offset
	end := paginator.Offset + paginator.Limit
	if start > count {
		start = count
	}
	if end > count {
		end = count
	}
	users := userRepo.users[start:end]

	return paginator.PagedData(users), nil
}

func (userRepo *userRepositoryImpl) Save(user *model.User) (*model.User, error) {
	userRepo.users = append(userRepo.users, user)
	return user, nil
}

func (userRepo *userRepositoryImpl) FindById(id string) (*model.User, error) {
	users := userRepo.users
	i, found := searchUserById(users, id)
	if !found {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}
	existUser := users[i]
	return existUser, nil
}

func (userRepo *userRepositoryImpl) Update(id string, user *model.User) (*model.User, error) {
	users := userRepo.users
	i, found := searchUserById(users, id)
	if !found {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}
	existUser := users[i]
	existUser.UserInput = user.UserInput
	users[i] = existUser

	return existUser, nil
}

func (userRepo *userRepositoryImpl) DeleteById(id string) error {
	users := userRepo.users
	i, found := searchUserById(users, id)
	if !found {
		return exception.ResourceNotFoundException("User", "id", id)
	}
	users = append(users[:i], users[i+1:]...)
	return nil
}

func searchUserById(users []*model.User, id string) (int, bool) {
	for i, v := range users {
		if v.ID == id {
			return i, true
		}
	}
	return -1, false
}
