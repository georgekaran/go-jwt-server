package repository

import (
	"database/sql"
)

type Repository interface {
	FindOne(id int)
	FindAll()
	Save(t interface{})
	SaveAll(t interface{})
	Delete(t interface{})
	DeleteAll(t interface{})
}

type BaseRepository struct {
	db *sql.DB
}