package router

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func HandlePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var newPerson Person

	requesterIP := r.RemoteAddr
	userAgent := r.UserAgent()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO person ( name, age) VALUES ( ?, ?)", newPerson.Name, newPerson.Age)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPerson)
	log.Printf("POST request processed successfully, from: %s with UA: %s", requesterIP, userAgent)

}
