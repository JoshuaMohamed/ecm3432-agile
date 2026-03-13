package persistence

import (
	"database/sql"
	"fmt"
	"log/slog"
	"server/logic"
	"strings"

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
	db, err := sql.Open("sqlite", "./server.db?_pragma=foreign_keys(1)&_pragma=busy_timeout(5000)")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	slog.Info("Connected to the SQLite database successfully")
	return db, nil
}

func (c *DatabaseClient) CreateTable(details logic.TableDetails) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", details.Name, details.Schema)
	_, err := c.db.Exec(query)
	return err
}

func (c *DatabaseClient) InsertRow(table string, fields []string, values []interface{}) error {
	placeholders := make([]string, len(fields))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
		table,
		strings.Join(fields, ","),
		strings.Join(placeholders, ","),
	)
	_, err := c.db.Exec(query, values...)
	return err
}

func (c *DatabaseClient) UpsertRow(table string, fields []string, values []interface{}) error {
	placeholders := make([]string, len(fields))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	updates := make([]string, len(fields))
	for i, f := range fields {
		updates[i] = fmt.Sprintf("%s = excluded.%s", f, f)
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ON CONFLICT DO UPDATE SET %s;",
		table,
		strings.Join(fields, ","),
		strings.Join(placeholders, ","),
		strings.Join(updates, ","),
	)
	_, err := c.db.Exec(query, values...)
	return err
}

func (c *DatabaseClient) DeleteRows(table, key, value string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?;", table, key)
	_, err := c.db.Exec(query, value)
	return err
}

func (c *DatabaseClient) GetPlaces(prefix string, limit, offset int) (logic.DBRows, error) {
	return c.db.Query("SELECT name,postcode,cover FROM Places WHERE postcode LIKE ? LIMIT ? OFFSET ?;", prefix+"%", limit, offset)
}

func (c *DatabaseClient) Query(table, key, value string) (logic.DBRows, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?;", table, key)
	return c.db.Query(query, value)
}

func (c *DatabaseClient) Close() error {
	return c.db.Close()
}
