package db

import (
	"database/sql"
	"fmt"
	"github.com/georgekaran/go-jwt-server/util"
	"github.com/georgekaran/go-jwt-server/util/file"
	_ "github.com/lib/pq"
)

var db *sql.DB
var dbError error

func GetConnection() *sql.DB {
	if db == nil {
		configMap := file.ToMap("config.properties")
		dataSourceName := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", configMap["database.user"],
			configMap["database.password"],
			configMap["database.host"],
			configMap["database.name"])
		db, dbError = sql.Open("postgres", dataSourceName)
		if dbError != nil {
			util.Must(dbError)
		}
	}
	return db
}

func CloseConnection() {
	db.Close()
}
