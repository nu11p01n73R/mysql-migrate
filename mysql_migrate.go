package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"syscall"
)

// Type connection defines a connection
// to the MySQL database.
type connection struct {
	Host     string
	Username string
	Dbname   string
	Password string
	Db       *sql.DB
}

// Connects to the database with the available
// credentials on conn
// returns
// 	-sql.DB on successsfull connection
//  -error on any error occured while connecting to db.
func (conn connection) Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", conn.Username, conn.Password, conn.Host, conn.Dbname)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, errors.New("Unable to open mysql connection")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.New("Database ping failed with error : " + err.Error())
	}
	return db, nil
}

// Parses command line flags
// and return an connection type
func parseConnectionFlags() *connection {
	// Available flags.
	host := flag.String("h", "", "MySQL host url")
	username := flag.String("u", "", "MySQL user name")
	dbname := flag.String("db", "", "MySQL database name")

	flag.Parse()

	var password string
	fmt.Printf("Enter password for %s@%s\n", *username, *host)
	bytePasssd, _ := terminal.ReadPassword(int(syscall.Stdin))
	password = string(bytePasssd)

	return &connection{Host: *host, Username: *username, Password: password, Dbname: *dbname}
}

func main() {
	connection := parseConnectionFlags()
	_, err := connection.Connect()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully connected")
}
