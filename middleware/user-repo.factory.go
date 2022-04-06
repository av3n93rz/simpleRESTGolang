package middleware

import (
	"database/sql"
	"net/http"
)

type UserRepo struct {
	CreateUser func(w http.ResponseWriter, req *http.Request)
	GetUser    func(w http.ResponseWriter, req *http.Request)
	GetAllUser func(w http.ResponseWriter, req *http.Request)
	DeleteUser func(w http.ResponseWriter, req *http.Request)
}

func UserRepoFactory(db *sql.DB) UserRepo {
	userRepo := UserRepo{
		CreateUser: CreateUser(db),
		GetUser:    GetUser(db),
		GetAllUser: GetUsers(db),
		DeleteUser: DeleteUser(db),
	}
	return userRepo
}
