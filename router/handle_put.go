package router

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func HandlePut(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var updateUser UpdatePerson

	requesterIP := r.RemoteAddr
	userAgent := r.UserAgent()

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE person SET name=?, age=? WHERE ID=?", updateUser.Name, updateUser.Age, updateUser.ID)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateUser)
	log.Printf("PUT request processed successfully, from: %s with UA: %s", requesterIP, userAgent)

}
