package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

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
func runCommand(command []string, conn Connection) error {
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
	_, err := getDbConnection()
	checkErrors(err)
}
