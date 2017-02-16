package main

import (
	"errors"
	"fmt"
)

func create(name string) error {
	err := createMigrationFile(name)
	return err
}

func migrate(conn Connection) error {
	if conn == (Connection{}) {
		return errors.New("No connection available to db")
	}

	err := createLogTable(conn)

	if err != nil {
		return err
	}

	applied, err := getAppliedMigrations(conn)
	if err != nil {
		return err
	}
	available, err := getAvailableMigrations()
	if err != nil {
		return err
	}

	fmt.Println(applied, available)
	return nil
}
