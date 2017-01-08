package main

import (
	"time"
)

type migration_log struct {
	Version    string
	Name       string
	ApplyTime  time.Time
	connection connection
}

// Checks if the migration_log table exists in the
// database connected by connection
// Returns
//	bool 	If the table exists or not
//	error 	Any error that occured
func (this *migration_log) CheckLogTable() (bool, error) {
	db, err := this.connection.Connect()
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

// Creates a new migration_log table.
// Return
//	bool 	If the table is created.
// 	error	Any error that occured
func (this *migration_log) CreateLogTable() (bool, error) {
	query := "CREATE TABLE `migration_log` (" +
		"`version` bigint(20) NOT NULL," +
		"`name` varchar(100) DEFAULT NULL," +
		"`apply_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"PRIMARY KEY (`version`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8;"

	db, err := this.connection.Connect()
	if err != nil {
		return false, err
	}

	statement, err := db.Prepare(query)
	if err != nil {
		return false, err
	}

	_, err = statement.Exec()
	if err != nil {
		return false, err
	}

	return true, err
}
