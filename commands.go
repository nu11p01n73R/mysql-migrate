package main

import (
	"errors"
	"flag"
	"fmt"
)

func create(name string) error {
	migration := Migration{
		Name: name}
	migration.GenerateVersion()
	err := migration.CreateFiles()
	return err
}

func migrate() error {
	return nil
}

func list() error {
	migrations, err := getAvailableMigrations()
	if err != nil {
		return err
	}

	fmt.Println(migrations)
	//available, err := getAvailableMigrations()
	//if err != nil {
	//		return err
	//	}
	//	fmt.Println(available)
	return nil
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
	case "create", "list":
		return command, nil
	default:
		return []string{}, errors.New("Unkown command supplied")
	}

}

// Runs a command
// Returns
// 	error Any error occured while running the command
func runCommand() error {
	var err error
	command, err := parseCommand()
	if err != nil {
		return err
	}

	switch command[0] {
	case "create":
		err = create(command[1])
	case "migrate":
		err = migrate()
	case "list":
		err = list()
	default:
		err = errors.New("Unknow command")
	}
	return err
}
