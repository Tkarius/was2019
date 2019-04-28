package main

import (
	"database/sql"
	"fmt"
	"log"
)

// global :s Let's see if we need to leave it that way.
// this is not exposed outside database.go.
var db *sql.DB

var dbServer = "jotain.database.windows.net"
var dbPort = 1433
var dbViewUser = "dbViewUuseri"
var dbViewUserSecret = "superSecure1"
var dbCreateUser = "dbCreateUuseri"
var dbCreateUserSecret = "superSecure2"
var database = "was-database"

func connectWithViewUser() {
	fmt.Printf("DEBUG: Connecting to db with view user.")
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		dbServer, dbViewUser, dbViewUserSecret, database)
	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

}

func connectWithCreateUser() {
	fmt.Printf("DEBUG: Connecting to db with create user.")
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		dbServer, dbCreateUser, dbCreateUserSecret, database)

}

func selectAnnouncements() {
	connectWithViewUser()
}
