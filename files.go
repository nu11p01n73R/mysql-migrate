package main

import (
	"os"
)

const MIGRATION_DIR = "./migrations"

type migration_file struct {
	FileName string
	Content  string
}

func checkMigrationDir() (bool, error) {
	if _, err := os.Stat(MIGRATION_DIR); os.IsNotExist(err) {
		return false, nil
	} else if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

func createMigrationDir() error {
	status, err := checkMigrationDir()

	if err != nil {
		return err
	}

	if !status {
		return os.Mkdir(MIGRATION_DIR, 0754)
	} else {
		return nil
	}
}
