package api

import (
	"encoding/json"
	"github.com/georgekaran/go-jwt-server/model"
	"github.com/georgekaran/go-jwt-server/service"
	"github.com/georgekaran/go-jwt-server/util"
	stringUtil "github.com/georgekaran/go-jwt-server/util/string"
	"net/http"
)

var userService service.UserService

func init() {
	userService = service.InitUserService()
}

func UserHandler(mux *http.ServeMux) {
	mux.HandleFunc("/api/user", handleIndex)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch {
	case r.Method == http.MethodPost:
		var user model.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if stringUtil.HasStringEmpty(user.Email, user.Password) {
			w.WriteHeader(http.StatusBadRequest)
			data, _ := json.Marshal(util.NewError("Missing fields."))
			w.Write(data)
		} else {
			saveError := userService.Save(user)
			if saveError != nil {
				w.WriteHeader(http.StatusBadRequest)
				data, _ := json.Marshal(util.NewError(saveError.Error()))
				w.Write(data)
			}
		}
	case r.Method == http.MethodGet:
		data, _ := json.Marshal(userService.FindAll())
		w.Write(data)
	}
}