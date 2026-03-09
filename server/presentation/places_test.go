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

func TestRouter_GetPlaces_Success(t *testing.T) {
	svc := &mockService{places: []logic.Place{{Name: "Exeter", Postcode: "EX4 4PY"}}}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/getPlaces?postcode=EX4%204PY&filter=district", nil)
	w := httptest.NewRecorder()

	rt.GetPlaces(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusOK)
	}
	var body presentation.GeneralResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	data, ok := body.Data.([]interface{})
	if !ok {
		t.Fatalf("expected Data to be []interface{}, got %T", body.Data)
	}
	if len(data) != 1 {
		t.Fatalf("data length = %d, want 1", len(data))
	}
}

func TestRouter_GetPlaces_DBError(t *testing.T) {
	svc := &mockService{err: errors.New("db down")}
	rt := presentation.NewRouter(svc)
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

func TestRouter_CreatePlace_Success(t *testing.T) {
	svc := &mockService{}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodPost, "/createPlace", strings.NewReader(`{"name":"The Tall Statue","postcode":"EX1 1AA","cover_path":"cover.png"}`))
	w := httptest.NewRecorder()

	rt.CreatePlace(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusOK)
	}
}

func TestRouter_CreatePlace_DBError(t *testing.T) {
	svc := &mockService{err: errors.New("db down")}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodPost, "/createPlace", strings.NewReader(`{"name":"The Tall Statue","postcode":"EX1 1AA","cover_path":"cover.png"}`))
	w := httptest.NewRecorder()

	rt.CreatePlace(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusInternalServerError)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if !strings.Contains(string(body), "Error: failed to create place") {
		t.Fatalf("unexpected response body: %s", string(body))
	}
}

func TestRouter_CreatePlace_BadRequest(t *testing.T) {
	svc := &mockService{}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodPost, "/createPlace", strings.NewReader(`not json`))
	w := httptest.NewRecorder()

	rt.CreatePlace(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusBadRequest)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if !strings.Contains(string(body), "Bad Request") {
		t.Fatalf("unexpected response body: %s", string(body))
	}
}
