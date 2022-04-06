package middleware

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"go-postgres/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUser(db *sql.DB) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		params := mux.Vars(req)

		id, err := strconv.Atoi(params["id"])

		if err != nil {
			log.Fatalf("Convertion error.  %v", err)
		}

		user, err := getUser(int64(id), db)

		if err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode("user not found")
		} else {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func getUser(id int64, db *sql.DB) (models.User, error) {
	var user models.User

	sqlStatement := `SELECT * FROM users WHERE id=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&user.ID, &user.Username, &user.Email)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("User not found")
		return user, errors.New("user not found")
	case nil:
		return user, nil
	default:
		log.Fatalf("Internal server error. %v", err)
		return user, err
	}
}
