package util

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) string {
	hashPassword, errorHashing := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errorHashing != nil {
		log.Println(errorHashing)
		return ""
	}
	return string(hashPassword)
}