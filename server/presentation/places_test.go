package presentation_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"server/logic"
	"server/presentation"
	"strings"
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

type getPlacesResponse struct {
	Data    []logic.Place `json:"data"`
	Message string        `json:"message"`
}

func TestRouter_GetPlaces_Success(t *testing.T) {
	db := mockDB{rows: &mockRows{rows: []logic.Place{{Name: "Exeter", Postcode: "EX4 4PY"}}}}
	rt := presentation.NewRouter(db)
	req := httptest.NewRequest(http.MethodGet, "/getPlaces?postcode=EX4%204PY&filter=district", nil)
	w := httptest.NewRecorder()

	rt.GetPlaces(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusOK)
	}
	var body getPlacesResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(body.Data) != 1 {
		t.Fatalf("data length = %d, want 1", len(body.Data))
	}
}

func TestRouter_GetPlaces_DBError(t *testing.T) {
	db := mockDB{getPlacesErr: errors.New("db down")}
	rt := presentation.NewRouter(db)
	req := httptest.NewRequest(http.MethodGet, "/getPlaces?postcode=EX4%204PY&filter=district", nil)
	w := httptest.NewRecorder()

	rt.GetPlaces(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusInternalServerError)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if !strings.Contains(string(body), "failed to get places") {
		t.Fatalf("unexpected response body: %s", string(body))
	}
}
