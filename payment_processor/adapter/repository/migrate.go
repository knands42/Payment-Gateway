package repository

import (
	"context"
	"database/sql"
	"io/fs"

	"github.com/maragudk/migrate"
	_ "github.com/mattn/go-sqlite3"
)

func Up(db *sql.DB, migrationsDir fs.FS) *sql.DB {
	if err := migrate.Up(context.Background(), db, migrationsDir); err != nil {
		panic(err)
	}
	return db
}

func Down(db *sql.DB, migrationsDir fs.FS) {
	if err := migrate.Down(context.Background(), db, migrationsDir); err != nil {
		panic(err)
	}
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
