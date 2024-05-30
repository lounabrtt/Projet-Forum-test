package function

import (
	"database/sql"
	// "fmt"
	// "log"
	"time"

	// _ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	BirthDate time.Time
	Email     string
	Password  string
}

func addUser(db *sql.DB, user User) error {
	// Préparation de la requête SQL
	query := "INSERT INTO users (first_name, last_name, birth_date, email, password) VALUES (?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Exécution de la requête avec les valeurs fournies
	_, err = stmt.Exec(user.FirstName, user.LastName, user.BirthDate, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil

	fmt.Fprintf(w, "User created successfully: %s\n", username)
    http.Redirect(w, r, "/connect", http.StatusFound)


}

