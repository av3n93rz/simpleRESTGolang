package middleware

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"go-postgres/models"
	"log"
	"net/http"
)

func GetUsers(db *sql.DB) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		users, err := getUsers(db)

		if err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(err.Error())
		} else {
			json.NewEncoder(w).Encode(users)
		}
	}
}

func getUsers(db *sql.DB) ([]models.User, error) {
	var users []models.User

	sqlStatement := `SELECT * FROM users`

	rows, err := db.Query(sqlStatement)
	defer rows.Close()

	noUsersError := errors.New("0 users were found")

	for rows.Next() {
		noUsersError = nil
		var user models.User
		err = rows.Scan(&user.ID, &user.Username, &user.Email)

		if err != nil {
			log.Fatalf("Internal server error. %v", err)
		}
		users = append(users, user)
	}

	if err != nil {
		log.Fatalf("Internal server error. %v", err)
		return users, err
	}

	if noUsersError != nil {
		fmt.Println(noUsersError.Error())
		return users, noUsersError
	}

	return users, nil

}
