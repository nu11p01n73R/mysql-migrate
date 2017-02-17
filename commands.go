package main

import (
	"fmt"
)

func create(name string) error {
	err := createMigrationFile(name)
	return err
}

func migrate() error {
	err := createLogTable()
	if err != nil {
		return err
	}

	applied, err := getAppliedMigrations()
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
