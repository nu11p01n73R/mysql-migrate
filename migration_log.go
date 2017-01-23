package main

import (
	"time"
)

const MIGRATION_TABLE = "migration_log"

type Migration struct {
	Version   string
	Name      string
	ApplyTime time.Time
}

// Checks if the migration_log table exists in the
// database connected by connection
// Returns
//	bool 	If the table exists or not
//	error 	Any error that occured
func checkLogTable(conn connection) (bool, error) {
	db, err := conn.Connect()
	if err != nil {
		return false, err
	}

	query := "SELECT 1 FROM migration_log LIMIT 1"
	_, err = db.Prepare(query)

	if err != nil {
		return false, nil
	}

	return true, nil
}

// Creates a new migration_log table,
// if not exists.
// Return
// 	error	Any error that occured
func createLogTable(conn connection) error {
	ok, err := checkLogTable(conn)
	if err != nil {
		return err
	}

	// If the table already exists
	if ok {
		return nil
	}

	query := "CREATE TABLE `migration_log` (" +
		"`version` bigint(20) NOT NULL," +
		"`name` varchar(100) DEFAULT NULL," +
		"`apply_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"PRIMARY KEY (`version`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8;"

	db, err := conn.Connect()
	if err != nil {
		return err
	}

	statement, err := db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil
}

func getAppliedMigrations(conn connection) ([]Migration, error) {
	migrations := []Migration{}

	query := "SELECT version, name, apply_time " +
		"FROM " + MIGRATION_TABLE

	db, err := conn.Connect()
	if err != nil {
		return []Migration{}, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return []Migration{}, err
	}

	for rows.Next() {
		migration := Migration{}
		var applyTime string

		err := rows.Scan(&migration.Version, &migration.Name, &applyTime)
		if err != nil {
			return []Migration{}, err
		}

		migration.ApplyTime, err = time.Parse("2006-01-02 15:04:05", applyTime)
		if err != nil {
			return []Migration{}, err
		}

		migrations = append(migrations, migration)
	}
	return migrations, nil

}
