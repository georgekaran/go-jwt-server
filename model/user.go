package model

import "github.com/georgekaran/go-jwt-server/util"

type User struct {
	BaseModel
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func NewUser(name, email, password string) User {
	user := User{Name: name, Email: email, Password: util.HashPassword(password)}
	return user
}