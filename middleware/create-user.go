package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-postgres/models"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func CreateUser(db *sql.DB) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var user models.User

		err := json.NewDecoder(req.Body).Decode(&user)

		if err != nil {
			log.Fatalf("Unable to decode the request body.  %v\n", err)
		}

		insertID := insertUser(user, db)

		res := response{
			ID:      insertID,
			Message: "User created successfully",
		}

		json.NewEncoder(w).Encode(res)
	}
}

func insertUser(user models.User, db *sql.DB) int64 {

	sqlStatement := `INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, user.Username, user.Email).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v\n", err)
	}

	fmt.Printf("Inserted a single record %v\n", id)

	return id
}
