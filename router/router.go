package router

import (
	"go-postgres/middleware"

	"github.com/gorilla/mux"
)

func Router(userRepo middleware.UserRepo) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/create", userRepo.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/delete/{id}", userRepo.DeleteUser).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/user/{id}", userRepo.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user", userRepo.GetAllUser).Methods("GET", "OPTIONS")

	return router
}
