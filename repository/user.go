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
	rows, err := u.db.Query("SELECT id, name, email FROM user_account")
	util.CheckFatal(err)

	users := make([]model.User, 0)
	for rows.Next() {
		user := model.User{}
		errScan := rows.Scan(&user.Id, &user.Name, &user.Email)
		util.CheckPrint(errScan)
		users = append(users, user)
	}

	return users
}

func (u UserRepository) Save(user model.User) error {
	// insert values
	_, err := u.db.Exec("INSERT INTO user_account (id, name, email, password) VALUES (DEFAULT, $1, $2, $3)", user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepository) Login(login, password string) (model.User, error) {
	var user model.User

	row := u.db.QueryRow("SELECT id, name, email, password FROM user_account WHERE email = $1 AND password = $2", login, password)

	errScan := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if errScan != nil {
		return user, errScan
	}
	return user, nil
}