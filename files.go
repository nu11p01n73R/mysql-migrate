package main

import (
	"os"
	"time"
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

func generateFileName(name string) string {
	// Reference date
	// Jan 2 15:04:05 2006 MST
	id := time.Now().Format("20060102150405")
	return MIGRATION_DIR + "/" + id + "_" + name + ".sql"
}

func createMigrationFile(name string) error {
	err := createMigrationDir()
	if err != nil {
		return err
	}

	fileName := generateFileName(name)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	err = file.Close()
	return err

}
