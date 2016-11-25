package main

import (
	"flag"
	"fmt"
)

// Type connection defines a connection
// to the MySQL database.
type connection struct {
	Host     string
	Username string
	Password string
}

// Parses command line flags
// and return an connection type
func parseConnectionFlags() *connection {
	// Available flags.
	host := flag.String("h", "", "MySQL host url")
	username := flag.String("u", "", "MySQL user name")
	password := flag.String("p", "", "MySQL password")

	flag.Parse()

	return &connection{Host: *host, Username: *username, Password: *password}
}

func main() {
	connection := parseConnectionFlags()
	fmt.Println(connection)
}
