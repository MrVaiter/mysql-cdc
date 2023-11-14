package cdc

import (
	"database/sql"
	"fmt"
)

func CheckGTID(db *sql.DB) (bool, error) {
	fmt.Print("\nChecking GTID mode...\n")

	rows, err := db.Query("SELECT @@global.gtid_mode;")
	if err != nil {
		return false, err
	}

	var result string
	for rows.Next() {
		err := rows.Scan(&result)

		if err != nil {
			return false, err
		}

		fmt.Print("GTID mode is " + result + "\n\n")

		if result == "OFF" {
			_, err := EnableGTIDMode(db)
			
			if err != nil {
				return false, err
			}
		}
	}

	return true, nil
}
