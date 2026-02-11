package logic

// PlacesRows abstraction to allow unit tests
type PlacesRows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Err() error
	Close() error
}

// Database defines the subset of database operations used by the logic layer.
type Database interface {
	CreateTable(details TableDetails) error
	GetPlaces(searchPrefix string, limit, offset int) (PlacesRows, error)
	Close() error
}

type TableDetails struct {
	Name   string
	Schema string
}
