package database

import (
	"database/sql"
	"flights-master/settings"
	"fmt"

	_ "embed"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func New(s *settings.Settings) (*sqlx.DB, error) {
	dbPath := fmt.Sprintf("file:%s", s.DbPath)

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	dbx := sqlx.NewDb(db, "sqlite")
	return dbx, nil
}

func PopulateDb(db *sqlx.DB) error {
	// _, err := db.Exec(schema)
	return nil
}
