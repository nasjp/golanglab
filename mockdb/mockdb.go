package mockdb

import "database/sql"

func QueryName(db *sql.DB) string {
	rows, err := db.Query("SELECT name FROM customers WHERE id=$1", 100)
	if err == sql.ErrNoRows {
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return ""
		}
		return name
	}

	if err := rows.Err(); err != nil {
		return ""
	}

	return ""
}
