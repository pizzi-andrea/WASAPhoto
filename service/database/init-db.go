package database

import (
	"bytes"
	"database/sql"
	"os"
	"path/filepath"
)

func initDb(db *sql.DB) (_error error) {
	path, _ := filepath.Abs("service/database/createTable.sql")

	byteQuery, _error := os.ReadFile(path)
	if _error != nil {
		return
	}

	query := bytes.NewBuffer(byteQuery).String()
	_, _error = db.Exec(query)

	return

}
