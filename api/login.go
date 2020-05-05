package api

import (
	"encoding/json"
	"github.com/georgekaran/go-jwt-server/auth"
	"github.com/georgekaran/go-jwt-server/service"
	"github.com/georgekaran/go-jwt-server/util"
	"log"
	"net/http"
)

type userCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type jwt struct {
	Token string `json:"token"`
}

func LoginHandler(mux *http.ServeMux) {
	mux.HandleFunc("/login", handleLogin)
}

func handleLogin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userCred userCredentials

	if req.Method == http.MethodPost {
		err := json.NewDecoder(req.Body).Decode(&userCred)
		if err != nil {
			http.Error(w, "Invalid fields.", http.StatusBadRequest)
			return
		}
		user, errLogin := service.UserServiceInstance.Login(userCred.Username, userCred.Password)
		if errLogin != nil {
			http.Error(w, "Invalid credentials.", http.StatusBadRequest)
		}
		token, errSign := auth.JWTInstance.Sign(user.Email)
		if errSign != nil {
			log.Fatal(errSign)
		}
		jwt := jwt{token}
		jwtJson, _ := json.Marshal(jwt)
		_, es := w.Write(jwtJson)
		util.Must(es)
	}
}