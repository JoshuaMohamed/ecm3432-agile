package logic_test

import (
	"errors"
	"reflect"
	"server/logic"
	"testing"
)

type mockDB struct {
	rows         logic.PlacesRows
	getPlacesErr error
}

func (m mockDB) CreateTable(details logic.TableDetails) error {
	return nil
}

func (m mockDB) GetPlaces(searchPrefix string, limit, offset int) (logic.PlacesRows, error) {
	if m.getPlacesErr != nil {
		return nil, m.getPlacesErr
	}
	return m.rows, nil
}

func (m mockDB) Close() error {
	return nil
}

type mockRows struct {
	rows []logic.Place
	idx  int
	err  error
}

func (m *mockRows) Next() bool {
	if m.err != nil {
		return false
	}
	if m.idx >= len(m.rows) {
		return false
	}
	m.idx++
	return true
}

func (m *mockRows) Scan(dest ...interface{}) error {
	if m.idx == 0 || m.idx > len(m.rows) {
		return errors.New("scan out of bounds")
	}
	if len(dest) != 2 {
		return errors.New("unexpected scan arg count")
	}
	namePtr, ok := dest[0].(*string)
	if !ok {
		return errors.New("invalid name dest")
	}
	postcodePtr, ok := dest[1].(*string)
	if !ok {
		return errors.New("invalid postcode dest")
	}
	row := m.rows[m.idx-1]
	*namePtr = row.Name
	*postcodePtr = row.Postcode
	return nil
}

func (m *mockRows) Err() error {
	return m.err
}

func (m *mockRows) Close() error {
	return nil
}

func TestGetPlaces(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		db       logic.Database
		postcode string
		filter   string
		limit    int
		offset   int
		wantErr  bool
	}{
		{
			name:     "success",
			postcode: "EX4 4PY",
			filter:   "district",
			limit:    10,
			offset:   0,
			wantErr:  false,
		},
		{
			name:     "invalid postcode",
			postcode: "BAD",
			filter:   "district",
			limit:    10,
			offset:   0,
			wantErr:  true,
		},
		{
			name:     "invalid filter",
			postcode: "EX4 4PY",
			filter:   "bad",
			limit:    10,
			offset:   0,
			wantErr:  true,
		},
		{
			name:     "db error",
			postcode: "EX4 4PY",
			filter:   "district",
			limit:    10,
			offset:   0,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "db error" {
				tt.db = mockDB{getPlacesErr: errors.New("db down")}
			} else if tt.name == "success" {
				tt.db = mockDB{rows: &mockRows{rows: []logic.Place{{Name: "Exeter", Postcode: "EX4 4PY"}}}}
			} else {
				tt.db = mockDB{}
			}
			got, gotErr := logic.GetPlaces(tt.db, tt.postcode, tt.filter, tt.limit, tt.offset)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetPlaces() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetPlaces() succeeded unexpectedly")
			}
			if tt.name == "success" {
				want := []logic.Place{{Name: "Exeter", Postcode: "EX4 4PY"}}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("GetPlaces() = %v, want %v", got, want)
				}
			}
		})
	}
}

func TestIsValidPostcode(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		postcode string
		want     bool
	}{
		{
			name:     "valid postcode",
			postcode: "EX4 4PY",
			want:     true,
		},
		{
			name:     "invalid postcode",
			postcode: "BAD",
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := logic.IsValidPostcode(tt.postcode)
			if got != tt.want {
				t.Errorf("IsValidPostcode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSearchPrefix(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		postcode string
		filter   string
		want     string
		wantErr  bool
	}{
		{
			name:     "area prefix",
			postcode: "EX4 4PY",
			filter:   "area",
			want:     "EX",
			wantErr:  false,
		},
		{
			name:     "invalid filter",
			postcode: "EX4 4PY",
			filter:   "bad",
			want:     "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := logic.GetSearchPrefix(tt.postcode, tt.filter)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetSearchPrefix() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetSearchPrefix() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("GetSearchPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
