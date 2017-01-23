package main

func create(name string) error {
	err := createMigrationFile(name)
	return err
}
