package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

// TestDB
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		fmt.Println(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err)

	}

	return db, func(tables ...string) {

		if len(tables) > 0 {
			db.Exec("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
		}

		db.Close()
	}
}
