package logic

import "database/sql"

// Database defines the subset of database operations used by the logic layer.
type Database interface {
	CreateTable(details TableDetails) error
	GetPlaces(searchPrefix string, limit, offset int) (*sql.Rows, error)
	Close() error
}

type TableDetails struct {
	Name   string
	Schema string
}
