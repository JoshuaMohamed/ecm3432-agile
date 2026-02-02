package persistence

import (
	"database/sql"
	"fmt"
	"log/slog"
	"server/logic"

	_ "github.com/glebarez/go-sqlite"
)

type DatabaseClient struct {
	db *sql.DB
}

func NewDatabaseClient() (*DatabaseClient, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	return &DatabaseClient{db: db}, nil
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./server.db?_pragma=foreign_keys(1)")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	slog.Info("Connected to the SQLite database successfully.")
	return db, nil
}

func (c *DatabaseClient) CreateTable(details logic.TableDetails) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", details.Name, details.Schema)
	_, err := c.Exec(query)
	return err
}

func (c *DatabaseClient) GetPlaces(prefix string, limit, offset int) (*sql.Rows, error) {
	return c.db.Query("SELECT name, postcode FROM Places WHERE postcode LIKE ? LIMIT ? OFFSET ?;", prefix+"%", limit, offset)
}

func (c *DatabaseClient) Exec(query string, args ...interface{}) (sql.Result, error) {
	return c.db.Exec(query, args...)
}

func (c *DatabaseClient) Close() error {
	return c.db.Close()
}
