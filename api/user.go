package api

import (
	"encoding/json"
	"github.com/georgekaran/go-jwt-server/model"
	"github.com/georgekaran/go-jwt-server/service"
	"github.com/georgekaran/go-jwt-server/util"
	stringUtil "github.com/georgekaran/go-jwt-server/util/string"
	"net/http"
)

const userPrefix string = "/api/user/"

var userService service.UserService

func init() {
	userService = service.InitUserService()
}

func UserHandler(mux *http.ServeMux) {
	mux.HandleFunc(userPrefix, handleIndex)
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch {
	case req.Method == http.MethodPost:
		values := req.PostForm
		email := values.Get("email")
		password := values.Get("password")
		if stringUtil.IsAtLeastOneStringEmpty(email, password) {
			w.WriteHeader(http.StatusBadRequest)
			data, _ := json.Marshal(util.NewError("Missing fields."))
			w.Write(data)
		}

		userService.Save(model.NewUser(email, password))
	case req.Method == http.MethodGet:
		data, _ := json.Marshal(userService.FindAll())
		w.Write(data)
	}
}