package main

import (
	"errors"
	"io/ioutil"
	"os"
	"regexp"
)

const MIGRATION_DIR = "./migrations"

// Check if migration directory exists or not
// Returns
//	bool of existance
//	error if any error occured in reading the directory.
func checkMigrationDir() (bool, error) {
	if _, err := os.Stat(MIGRATION_DIR); os.IsNotExist(err) {
		return false, nil
	} else if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

// Creates the migration directory.
// Returns
//	error, if some error occured in creating the directory.
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

// Parse an input file name to obtain the version and
// name of the migration.
// Returns
//	version of migrtion
//	name of migration
//	error If the filename doesn't followe the pattern or
//	 error in compiling regex
func parseFileNames(fileName string) (string, string, error) {
	regex, err := regexp.Compile("^([0-9]{14})_(.*?)_(?:up|down)\\.sql$")
	if err != nil {
		return "", "", err
	}
	parts := regex.FindStringSubmatch(fileName)
	if len(parts) != 3 {
		return "", "", errors.New("Unknown file format")
	}
	return parts[1], parts[2], nil

}

// List the migration files available in the
// migration directory.
// Returns
//	map of version to its Migration object
// 	error any error occured while reading the
//	 files
func getAvailableMigrations() (map[string]Migration, error) {
	migrations := map[string]Migration{}

	files, err := ioutil.ReadDir(MIGRATION_DIR)
	if err != nil {
	}
	for _, file := range files {
		version, name, err := parseFileNames(file.Name())
		if err != nil {
			// TODO maybe we can silently ignore this message.
			return migrations, err
		}

		_, ok := migrations[version]
		if ok {
			// up or down migration file of this
			// is already read. Skip this file
			continue
		}

		migrations[version] = Migration{
			Name:    name,
			Version: version}
	}
	return migrations, nil
}

// Creates a migration file.
// Returns
// 	string full path to the new file
//	error if any error occured while creating the file
func createMigrationFile(fileName string) (string, error) {
	fileName = MIGRATION_DIR + "/" + fileName
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	if err = file.Close(); err != nil {
		return "", err
	}
	return fileName, nil
}
