package router

import (
	"fmt"
	"io"
	"net/http"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	const apiUrl = "https://reqres.in/api/users/1"

	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Printf("Error GET Request: %v\n", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Statuscode: %v\n", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error handle: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func HandleRouting(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		fmt.Fprint(w, "INI POST")
	default:
		http.Error(w, "Method now allowed", http.StatusMethodNotAllowed)
	}
}
