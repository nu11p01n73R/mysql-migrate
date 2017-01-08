package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"syscall"
)

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

	return &connection{Host: *host, Username: *username,
		Password: password, Dbname: *dbname}
}

func checkErrors(err) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	connection := parseConnectionFlags()
	_, err := connection.Connect()
	checkErrors(err)

	_, err = connection.Connect()
	checkErrors(err)
}
