/*

The package works on 2 tables on a PostgreSQL data base server.

The names of the tables are:

	* Users
	* Userdata

The definitions of the tables in the PostgreSQL server are:

	CREATE TABLE Users (
    	ID SERIAL,
    	Username VARCHAR(100) PRIMARY KEY
	);

	CREATE TABLE Userdata (
    	UserID Int NOT NULL,
    	Name VARCHAR(100),
    	Surname VARCHAR(100),
    	Description VARCHAR(200)
	);

	This is rendered as code

This is not rendered as code

*/
package document

// BUG(1): Function ListUsers() not working as expected
// BUG(2): Function AddUser() is too slow

import (
	"database/sql"
	"fmt"
	"strings"
)

/*
This block of global variables holds the connection details to the Postgres server
	Hostname: is the IP or the hostname of the server
 	Port: is the TCP port the DB server listens to
	Username: is the username of the database user
	Password: is the password of the database user
	Database: is the name of the Database in PostgreSQL
*/
var (
	Hostname = ""
	Port     = 2345
	Username = ""
	Password = ""
	Database = ""
)

// The Userdata structure is for holding full user data
// from the Userdata table and the Username from the
// Users table
type Userdata struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Description string
}

// openConnection() is for opening the Postgres connection
// in order to be used by the other functions of the package.
func openConnection() (*sql.DB, error) {
	var db *sql.DB
	return db, nil
}

// The function returns the User ID of the username
// -1 if the user does not exist
func exists(username string) int {
	fmt.Println("Searching user", username)
	return 0
}

// AddUser adds a new user to the database
//
// Returns new User ID
// -1 if there was an error
func AddUser(d Userdata) int {
	d.Username = strings.ToLower(d.Username)
	return -1
}

/*
	DeleteUser deletes an existing user if the user exists.

	It requires the User ID of the user to be deleted.
*/
func DeleteUser(id int) error {
	fmt.Println(id)
	return nil
}

// ListUsers lists all users in the database
// and returns a slice of Userdata.
func ListUsers() ([]Userdata, error) {
	// Data holds the records returned by the SQL query
	Data := []Userdata{}
	return Data, nil
}

// UpdateUser is for updating an existing user
// given a Userdata structure.
// The user ID of the user to be updated is found
// inside the function.
func UpdateUser(d Userdata) error {
	fmt.Println(d)
	return nil
}
