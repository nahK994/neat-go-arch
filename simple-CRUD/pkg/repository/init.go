package repository

import (
	"database/sql"
	"fmt"
	"simple-CRUD/pkg/app"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func createTables(db *sql.DB) error {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(20) NOT NULL,
		email VARCHAR(50) NOT NULL UNIQUE,
		age INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	if _, err := db.Exec(createUserTable); err != nil {
		return fmt.Errorf("failed to create 'users' table: %w", err)
	}

	return nil
}

func Init(dbConfig *app.DB_config) (*Repository, error) {
	var err error
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Domain, dbConfig.Port, dbConfig.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	if err := createTables(db); err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}
