package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// User represents a user in the database
type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

// CreateUser inserts a new user into the database
func CreateUser(db *sql.DB, user User) error {
	query := "INSERT INTO users (name, email, created_at) VALUES (?, ?, ?)"
	_, err := db.Exec(query, user.Name, user.Email, user.CreatedAt)
	return err
}

// GetUser retrieves a user from the database by ID
func GetUser(db *sql.DB, id int) (User, error) {
	query := "SELECT id, name, email, created_at FROM users WHERE id = ?"
	row := db.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	return user, err
}

// UpdateUser updates a user's details in the database
func UpdateUser(db *sql.DB, user User) error {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := db.Exec(query, user.Name, user.Email, user.ID)
	return err
}

// DeleteUser deletes a user from the database by ID
func DeleteUser(db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, id)
	return err
}

func main() {
	// Example usage (make sure to replace with your actual database credentials)
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
