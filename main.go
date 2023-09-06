package main

import (
	"belajar/kedua/cmd/router"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/project_golang_satu")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to Database")
	return db, nil
}

func handleRoute(db *sql.DB) {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		router.GetItems(w, r, db)
	})

	http.HandleFunc("/users/add", func(w http.ResponseWriter, r *http.Request) {
		router.HandlePost(w, r, db)
	})

	http.HandleFunc("/users/edit", func(w http.ResponseWriter, r *http.Request) {
		router.HandlePut(w, r, db)
	})

	http.HandleFunc("/users/delete", func(w http.ResponseWriter, r *http.Request) {
		router.HandleDelete(w, r, db)
	})
}

func main() {

	//init DB
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//setup handler
	handleRoute(db)

	//start server
	fmt.Println("Server is running on :8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
