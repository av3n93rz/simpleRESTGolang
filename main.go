package main

import (
	"fmt"
	"go-postgres/middleware"
	"go-postgres/router"
	"log"
	"net/http"
)

func main() {
	db := middleware.ConnectToDb()
	defer db.Close()

	userRepo := middleware.UserRepoFactory(db)

	router := router.Router(userRepo)

	fmt.Println("Starting server on port 4200...")

	log.Fatal(http.ListenAndServe(":4200", router))
}
