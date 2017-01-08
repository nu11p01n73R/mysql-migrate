package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
func (conn *connection) Connect() (*sql.DB, error) {

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
