package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteUser(db *sql.DB) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		params := mux.Vars(req)

		id, err := strconv.Atoi(params["id"])

		if err != nil {
			log.Fatalf("Convertion error.  %v", err)
		}

		deletedRows := deleteUser(int64(id), db)

		msg := fmt.Sprintf("User has been deleted. Total rows affected %v", deletedRows)

		res := response{
			ID:      int64(id),
			Message: msg,
		}

		json.NewEncoder(w).Encode(res)
	}
}

func deleteUser(id int64, db *sql.DB) int64 {
	sqlStatement := `DELETE FROM users WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Something went wrong. %v\n", err)
	}

	r, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v\n", err)
	}

	fmt.Printf("Total rows affected %v\n", r)
	fmt.Println(r)
	return r
}
