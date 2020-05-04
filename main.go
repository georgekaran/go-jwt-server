package main

import (
	"github.com/georgekaran/go-jwt-server/api"
	"github.com/georgekaran/go-jwt-server/config"
	"github.com/georgekaran/go-jwt-server/db"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	api.UserHandler(mux)
	api.LoginHandler(mux)

	log.Printf("Running on port %s\n", config.ConfigMap["server.port"])
	err := http.ListenAndServe(config.ConfigMap["server.port"], mux)
	if err != nil {
		log.Fatal(err)
	}

	defer db.CloseConnection()
}