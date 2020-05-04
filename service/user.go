package service

import (
	"github.com/georgekaran/go-jwt-server/model"
	"github.com/georgekaran/go-jwt-server/repository"
)

type UserService struct {
	Repository repository.UserRepository
}

var UserServiceInstance UserService

func init() {
	UserServiceInstance = InitUserService()
}

func InitUserService() UserService {
	userService := UserService{
		Repository: repository.InitRepository(),
	}
	return userService
}

func (us *UserService) FindAll() []model.User {
	return UserServiceInstance.Repository.FindAll()
}

func (us *UserService) Save(user model.User) error {
	return UserServiceInstance.Repository.Save(user)
}

func (us *UserService) Login(login, password string) (model.User, error) {
	return UserServiceInstance.Repository.Login(login, password)
}