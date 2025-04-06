package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"simple-CRUD/pkg/app"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

func MigrateDB(dbConfig *app.DB_config) (*Repository, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Domain, dbConfig.Port, dbConfig.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	dir, _ := os.Getwd()
	migrationDir := filepath.Join(dir, "..", "pkg", "repository", "migrations")

	if err := goose.Up(db, migrationDir); err != nil {
		log.Fatal(err)
	}

	return &Repository{
		db: db,
	}, nil
}
