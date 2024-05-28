package main

import (
	"database/sql"
	"fmt"
	"net/http"

	 _ "github.com/mattn/go-sqlite3"
)

// Define the port variable
const port = ":8080"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/News.html")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/form.html")
}

func form2Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/form2.html")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bjr"))
}

func StaticFiles(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
}


func main() {
	// Database connection
	db, err := sql.Open("sqlite3", "database/forum.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Create the tables
	if err := CreateTableUser(db); err != nil {
		fmt.Println("Error creating user table:", err)
		return
	}
	if err := CreateTableCategories(db); err != nil {
		fmt.Println("Error creating categories table:", err)
		return
	}
	if err := CreateTablePost(db); err != nil {
		fmt.Println("Error creating post table:", err)
		return
	}

	// HTTP Handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/news", newsHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/home", indexHandler)
	http.HandleFunc("/css/", StaticFiles)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/form2", form2Handler)


	// Start the server
	fmt.Println("\n(http://localhost:8080/home) - Server started on port", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func CreateTableUser(db *sql.DB) error {
	// Creating the user table if not already created
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username VARCHAR(12) NOT NULL,
            password VARCHAR(12) NOT NULL,
			email TEXT NOT NULL,
			isAdmin BOOL NOT NULL DEFAULT FALSE,
			isBanned BOOL NOT NULL DEFAULT FALSE,
			pp BLOB,
			UUID VARCHAR(36) NOT NULL
        )
    `)
	return err
}

func CreateTableCategories(db *sql.DB) error {
	// Creating the categories table if not already created
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) NOT NULL UNIQUE,
			number_of_posts INTEGER DEFAULT 0
		)
	`)
	return err
}

func CreateTablePost(db *sql.DB) error {
	// Creating the post table if not already created
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(100) NOT NULL,
			content TEXT NOT NULL,
			date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			category INTEGER NOT NULL,
			author INTEGER NOT NULL,
			FOREIGN KEY(category) REFERENCES categories(id),
			FOREIGN KEY(author) REFERENCES users(id)
		)
	`)
	return err
}