package logic

// Rows abstraction to allow unit tests
type DBRows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Err() error
	Close() error
}

// Database defines the subset of database operations used by the logic layer.
type Database interface {
	CreateTable(details TableDetails) error
	InsertRow(table string, fields []string, values []interface{}) error
	UpsertRow(table string, fields []string, values []interface{}) error
	GetPlaces(searchPrefix string, limit, offset int) (DBRows, error)
	Close() error
}

type TableDetails struct {
	Name   string
	Schema string
}
