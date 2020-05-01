package model

import "github.com/georgekaran/go-jwt-server/util"

type User struct {
	BaseModel
	Email string `json:"email"`
	Password string `json:"password"`
}

func NewUser(email, password string) User {
	user := User{Email: email, Password: util.HashPassword(password)}
	return user
}