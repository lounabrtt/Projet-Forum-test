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
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Create the tables
	err = CreateTable(db)
	if err != nil {
		panic(err.Error())
	}

	// HTTP Handlers
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/news", newsHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/home", indexHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/form2", form2Handler)
	http.HandleFunc("/css/", StaticFiles)
	

	// // Static files handlers
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))
	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("."))))
	// http.Handle("/pictures/", http.StripPrefix("/pictures/", http.FileServer(http.Dir("."))))

	// Start the server
	fmt.Println("\n(http://localhost:8080/home) - Server started on port", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		panic(err.Error())
	}
}

// Function that creates a table User
func CreateTableUser(db *sql.DB) {
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
	if err != nil {
		panic(err.Error())
	}
}

// Function that creates a table Categories
func CreateTableCategories(db *sql.DB) {
	// Creating the categories table if not already created
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) NOT NULL UNIQUE,
			number_of_posts INTEGER DEFAULT 0
		)
	`)
	if err != nil {
		panic(err.Error())
	}
}

// Function that creates a table Post
func CreateTablePost(db *sql.DB) {
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
	if err != nil {
		panic(err.Error())
	}
}

// Function that creates all necessary tables
func CreateTable(db *sql.DB) error {
	CreateTableUser(db)
	CreateTableCategories(db)
	CreateTablePost(db)
	return nil
}
