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
