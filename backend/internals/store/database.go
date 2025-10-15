package store

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

func Open() (*sql.DB, error) {
	
	// * For Local development
	// db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")

	// Using NeonDB
	DbUrl := os.Getenv("DATABASE_URL")

	db, err := sql.Open("pgx", DbUrl)
	if err != nil {
		return nil, fmt.Errorf("DB Open: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Database Ping error: %w", err)
	}

	fmt.Println("Database Connected !!")
	return db, nil
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("Migrate : %w", err)
	}

	if err = goose.Up(db, dir); err != nil {
		return fmt.Errorf("Goose Up error: %w", err)
	}
	return nil
}

func MigrateFs(db *sql.DB, dir string, MigrationFS fs.FS) error {
	goose.SetBaseFS(MigrationFS)

	defer func() {
		goose.SetBaseFS(nil)
	}()

	return Migrate(db, dir)
}
