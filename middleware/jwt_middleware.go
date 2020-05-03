package middleware

import (
	"github.com/georgekaran/go-jwt-server/auth"
	"github.com/georgekaran/go-jwt-server/util/file"
	"net/http"
	"strings"
)

var jwt auth.JWT

func init() {
	configMap := file.ToMap("config.properties")
	jwt = auth.JWT{
		Secret: []byte(configMap["jwt.secret"]),
	}
}

func ValidateJwt(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		tokenParts := strings.Split(token, " ")
		if tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}
		_, tokenError := jwt.ValidateToken(tokenParts[1])
		if tokenError != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		jwt.ValidateToken(tokenParts[1])
		next(w, r)
	})
}