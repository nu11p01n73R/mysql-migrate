package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

// Type connection defines a connection
// to the MySQL database.
type Connection struct {
	Host     string
	Username string
	Dbname   string
	Password string
	Db       *sql.DB
}

// Singleton connection object.
var connection Connection

// Connects to the database with the available
// credentials on conn
// returns
// 	-sql.DB on successsfull connection
//  -error on any error occured while connecting to db.
func (conn *Connection) Connect() (*sql.DB, error) {

	if conn.Db != nil {
		return conn.Db, nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", conn.Username,
		conn.Password, conn.Host, conn.Dbname)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, errors.New("Unable to open mysql connection")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.New("Database ping failed with error : " +
			err.Error())
	}

	conn.Db = db
	return db, nil
}

// Parses command line flags
// and return an connection type
func parseConnectionFlags() (Connection, error) {
	// Available flags.
	host := flag.String("h", "localhost:3306", "MySQL host url")
	username := flag.String("u", "", "MySQL user name")
	dbname := flag.String("db", "", "MySQL database name")

	flag.Parse()

	if *username == "" {
		return Connection{}, errors.New("User name cannot be empty")
	}
	if *dbname == "" {
		return Connection{}, errors.New("Database name cannot be empty")
	}

	var password string
	fmt.Printf("Enter password for %s@%s\n", *username, *host)
	bytePasssd, _ := terminal.ReadPassword(int(syscall.Stdin))
	password = string(bytePasssd)

	return Connection{Host: *host, Username: *username,
		Password: password, Dbname: *dbname}, nil
}

// Single source for obtaining a database connection.
func getDbConnection() (*sql.DB, error) {
	// If already connected, return the connection.
	if connection.Db != nil {
		return connection.Db, nil
	}

	// Required to prevent declaring new local connection
	var err error
	connection, err = parseConnectionFlags()
	if err != nil {
		return nil, err
	}

	_, err = connection.Connect()
	if err != nil {
		return nil, err
	}

	return connection.Db, nil
}
