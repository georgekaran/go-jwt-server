package api

import (
	"encoding/json"
	"github.com/georgekaran/go-jwt-server/auth"
	"github.com/georgekaran/go-jwt-server/service"
	"log"
	"net/http"
)

type userCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(mux *http.ServeMux) {
	mux.HandleFunc("/login", handleLogin)
}

func handleLogin(w http.ResponseWriter, req *http.Request) {
	var userCred userCredentials

	if req.Method == http.MethodPost {
		err := json.NewDecoder(req.Body).Decode(&userCred)
		if err != nil {
			http.Error(w, "Invalid fields.", http.StatusBadRequest)
			return
		}
		_, errLogin := service.UserServiceInstance.Login(userCred.Username, userCred.Password)
		if errLogin != nil {
			http.Error(w, "Invalid credentials.", http.StatusBadRequest)
		}
		token, errSign := auth.JWTInstance.Sign("user")
		if errSign != nil {
			log.Fatal(errSign)
		}
		w.Write([]byte(token))
	}
}