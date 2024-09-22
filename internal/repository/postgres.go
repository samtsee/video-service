package repository

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo() (*PostgresRepo, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	user := os.Getenv("PSQL_USER")
	password := os.Getenv("PSQL_PASSWORD")
	port := os.Getenv("PSQL_PORT")
	dbName := os.Getenv("PSQL_DB_NAME")
	connStr := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", user, password, port, dbName)

	repo := &PostgresRepo{}
	if err := repo.Connect(connStr); err != nil {
		return nil, err
	}

	return repo, nil
}

func (p *PostgresRepo) Connect(connectionString string) error {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	p.db = db

	err = p.db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

func (p *PostgresRepo) Disconnect() error {
	err := p.db.Close()
	if err != nil {
		return fmt.Errorf("failed to close postgres connection: %w", err)
	}
	return nil
}
