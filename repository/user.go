package repository

import (
	"github.com/georgekaran/go-jwt-server/db"
	"github.com/georgekaran/go-jwt-server/model"
	"github.com/georgekaran/go-jwt-server/util"
)

type UserRepository BaseRepository

func InitRepository() UserRepository {
	br := UserRepository{db: db.GetConnection()}
	return br
}

func (u UserRepository) FindAll() []model.User {
	rows, err := u.db.Query("SELECT * FROM user_account")
	util.CheckFatal(err)

	users := make([]model.User, 0)
	for rows.Next() {
		user := model.User{}
		errScan := rows.Scan(&user.Id, &user.Email, &user.Password)
		util.CheckPrint(errScan)
		users = append(users, user)
	}

	return users
}