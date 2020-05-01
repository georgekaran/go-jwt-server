package util

import "log"

type Error struct {
	Message string `json:"message"`
}

func NewError(message string) Error {
	return Error{message}
}

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckPrint(err error) {
	if err != nil {
		log.Println(err)
	}
}