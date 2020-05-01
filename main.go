package main

import (
	"github.com/georgekaran/go-jwt-server/api"
	"github.com/georgekaran/go-jwt-server/db"
	"github.com/georgekaran/go-jwt-server/util/file"
	"log"
	"net/http"
)

var ConfigMap map[string]string

func init() {
	ConfigMap = file.ToMap("config.properties")
}

func main() {
	mux := http.NewServeMux()
	api.UserHandler(mux)

	log.Printf("Running on port %s\n", ConfigMap["server.port"])
	err := http.ListenAndServe(ConfigMap["server.port"], mux)
	if err != nil {
		log.Fatal(err)
	}

	defer db.CloseConnection()
}