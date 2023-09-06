package router

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func HandleDelete(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var deleteRequest struct {
		ID int `json:"ID"`
	}

	requesterIP := r.RemoteAddr
	userAgent := r.UserAgent()

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&deleteRequest)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM person WHERE ID=?", deleteRequest.ID)
	if err != nil {
		log.Println("Error delete record: ", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Record deleted success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(response)
	log.Printf("DELETE request processed successfully, from: %s with UA: %s", requesterIP, userAgent)

}
