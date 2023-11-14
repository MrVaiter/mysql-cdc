package cdc

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func GetGTIDs(config *Config) ([]string, error) {
	// Відкриваємо з'єднання з базою даних відповідно до конфігу
	db, err := sql.Open(
		config.Flavor,
		config.User+":"+
			config.Password+"@tcp("+
			config.Host+":"+
			strconv.Itoa(config.Port)+")/"+
			config.DB)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Перевірка статусу GTID режиму
	CheckGTID(db)

	// Виконання SQL-запиту для отримання GTID
	rows, err := db.Query("SELECT @@global.gtid_executed;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Читання результатів ключів транзакцій
	result := make([]string, 0, 3)
	var column string
	for rows.Next() {
		if err := rows.Scan(&column); err != nil {
			return nil, err
		}

		result = append(result, column)
	}

	// Перевірка на помилку
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
