package main

import (
	"errors"
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
	host := flag.String("h", "localhost:3306", "MySQL host url")
	username := flag.String("u", "", "MySQL user name")
	dbname := flag.String("db", "", "MySQL database name")

	flag.Parse()

	var password string
	if *username != "" && *dbname != "" {
		fmt.Printf("Enter password for %s@%s\n", *username, *host)
		bytePasssd, _ := terminal.ReadPassword(int(syscall.Stdin))
		password = string(bytePasssd)

		return &connection{Host: *host, Username: *username,
			Password: password, Dbname: *dbname}
	}

	return &connection{}

}

// Parses command line arguments to obtain the command
// to be executed.
// Returns
// 	[]string command followed by option
//	error If the command is not known or invalid number
func parseCommand() ([]string, error) {
	flag.Parse()
	command := flag.Args()

	if !(len(command) == 1 || len(command) == 2) {
		return []string{}, errors.New("Invalid number of parameters")
	}

	switch command[0] {
	case "create", "migrate", "rollback":
		return command, nil
	default:
		return []string{}, errors.New("Unkown command supplied")
	}

}

// Runs a command
// Returns
// 	error Any error occured while running the command
func runCommand(command []string, conn connection) error {
	var err error
	switch command[0] {
	case "create":
		err = create(command[1])
	case "migrate":
		err = migrate(conn)
	default:
		err = errors.New("Unknow command")
	}
	return err
}

// Handle errors.
func checkErrors(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	conn := parseConnectionFlags()
	cmd, err := parseCommand()
	checkErrors(err)

	err = runCommand(cmd, *conn)
	checkErrors(err)
}
