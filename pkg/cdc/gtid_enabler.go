package cdc

import (
	"database/sql"
	"fmt"
)

func EnableGTIDMode(db *sql.DB) (bool, error) {
	fmt.Print("\nChanging GTID mode...\n")

	rows, err := db.Query("SET @@GLOBAL.ENFORCE_GTID_CONSISTENCY = WARN;")
	if err != nil {
		return false, err
	}
	defer rows.Close()

	rows, err = db.Query("SET @@GLOBAL.ENFORCE_GTID_CONSISTENCY = ON;")
	if err != nil {
		return false, err
	}

	rows, err = db.Query("SET @@GLOBAL.GTID_MODE = OFF_PERMISSIVE;")
	if err != nil {
		return false, err
	}

	rows, err = db.Query("SET @@GLOBAL.GTID_MODE = ON_PERMISSIVE;")
	if err != nil {
		return false, err
	}

	rows, err = db.Query("SET @@GLOBAL.GTID_MODE = ON;")
	if err != nil {
		return false, err
	}

	fmt.Print("GTID mode enabled!\n\n")

	return true, nil
}
