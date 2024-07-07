package main

import (
	"database/sql"
	"testing"
	"time"
	"log"
	"regexp"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"os"
	"github.com/stretchr/testify/assert"
)

// Global variables for the mock database and sqlmock
var (
	db   *sql.DB
	mock sqlmock.Sqlmock
	err  error
)

// MainTest sets up the global sqlmock instance
func TestMain(m *testing.M) {
	fmt.Println("testing start")
	fmt.Println("============================")

	var dbMock *sql.DB
	dbMock, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db = dbMock
	defer db.Close()

	code := m.Run()

	fmt.Println("============================")
	fmt.Println("testing end")
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	// Define the expected SQL query and arguments
	query := regexp.QuoteMeta("INSERT INTO users (name, email, created_at) VALUES (?, ?, ?)")
	mock.ExpectExec(query).WithArgs("John Doe", "john@example.com", sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

	// Create a new user
	user := User{
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
	}

	// Call the CreateUser function and assert no error is returned
	err := CreateUser(db, user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUser(t *testing.T) {
	// Define the expected SQL query and result rows
	query := regexp.QuoteMeta("SELECT id, name, email, created_at FROM users WHERE id = ?")
	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at"}).AddRow(1, "John Doe", "john@example.com", time.Now())
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	// Call the GetUser function
	user, err := GetUser(db, 1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "John Doe", user.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateUser(t *testing.T) {
	// Define the expected SQL query and arguments
	query := regexp.QuoteMeta("UPDATE users SET name = ?, email = ? WHERE id = ?")
	mock.ExpectExec(query).WithArgs("Jane Doe", "jane@example.com", 1).WillReturnResult(sqlmock.NewResult(1, 1))

	// Create a user object with the updated details
	user := User{
		ID:    1,
		Name:  "Jane Doe",
		Email: "jane@example.com",
	}

	// Call the UpdateUser function and assert no error is returned
	err := UpdateUser(db, user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteUser(t *testing.T) {
	// Define the expected SQL query and arguments
	query := regexp.QuoteMeta("DELETE FROM users WHERE id = ?")
	mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the DeleteUser function and assert no error is returned
	err := DeleteUser(db, 1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
