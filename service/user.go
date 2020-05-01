package service

import (
	"fmt"
	"github.com/georgekaran/go-jwt-server/model"
	"github.com/georgekaran/go-jwt-server/repository"
)

type UserService struct {
	Repository repository.UserRepository
}

var userService UserService

func InitUserService() UserService {
	userService = UserService{
		Repository: repository.InitRepository(),
	}
	return userService
}

func (us *UserService) FindAll() []model.User {
	return userService.Repository.FindAll()
}

func (us *UserService) Save(u model.User) {
	fmt.Println("Estou salvando...", u)
}

func (UserService) SaveAll() {

}
