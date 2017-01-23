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

func generateFileName(name string) []string {
	// Reference date
	// Jan 2 15:04:05 2006 MST
	id := time.Now().Format("20060102150405")
	base := MIGRATION_DIR + "/" + id + "_" + name
	return []string{
		base + "_up.sql",
		base + "_down.sql"}
}

func createMigrationFile(name string) error {
	err := createMigrationDir()
	if err != nil {
		return err
	}

	fileNames := generateFileName(name)
	for _, fileName := range fileNames {
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0666)
		if err != nil {
			return err
		}

		if err = file.Close(); err != nil {
			return err
		}
	}
	return nil
}
