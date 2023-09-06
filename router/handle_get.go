package router

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func GetItems(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var persons []UpdatePerson

	requesterIP := r.RemoteAddr
	userAgent := r.UserAgent()

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		log.Println(err)
		http.Error(w, "Database Error: ", http.StatusInternalServerError)
	}
	defer rows.Close()

	for rows.Next() {
		var person UpdatePerson
		if err := rows.Scan(&person.ID, &person.Name, &person.Age); err != nil {
			log.Println(err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		persons = append(persons, person)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(persons); err != nil {
		log.Fatal(err)
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
		return
	}

	log.Printf("GET request processed successfully, from: %s with UA: %s", requesterIP, userAgent)
}
