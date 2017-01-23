package main

import (
	"errors"
	"fmt"
)

func create(name string) error {
	err := createMigrationFile(name)
	return err
}

func migrate(conn connection) error {
	if conn == (connection{}) {
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

	fmt.Println(applied)
	return nil
}
