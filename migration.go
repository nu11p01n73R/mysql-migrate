package main

import (
	"fmt"
	"time"
)

type Migration struct {
	Name      string
	Version   string
	ApplyTime time.Time
}

// Generate id for a new migration
func (this *Migration) GenerateVersion() {
	// Reference date
	// Jan 2 15:04:05 2006 MST
	this.Version = time.Now().Format("20060102150405")
}

// Get up and down file names for a particular
// migrations.
// The filenames follow pattern,
//	20060102150405_Name_up.sql
//	20060102150405_Name_down.sql
//
// Returns
//	Array of up and down filenames
func (this Migration) GetFileNames() []string {
	base := this.Version + "_" + this.Name
	return []string{
		base + "_up.sql",
		base + "_down.sql"}
}

// Creates the up and down migration files in the
// MIGRATION_DIR
// Function prints the path to the files that
// are created.
// Returns
//	error if some error occured in creating files.
func (this Migration) CreateFiles() error {
	err := createMigrationDir()
	if err != nil {
		return err
	}

	fmt.Println("Creating migration files:")
	fileNames := this.GetFileNames()
	for _, fileName := range fileNames {
		path, err := createMigrationFile(fileName)
		if err != nil {
			return err
		}
		fmt.Println("\t", path)
	}
	return nil
}
