package presentation_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"server/presentation"
	"strings"
	"testing"
)

func TestRouter_CreateAccount_Success(t *testing.T) {
	svc := &mockService{}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodPost, "/createAccount", strings.NewReader(`{"email":"user@example.com","password":"secret","role":"tourist"}`))
	w := httptest.NewRecorder()

	rt.CreateAccount(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusOK)
	}
}

func TestRouter_CreateAccount_DBError(t *testing.T) {
	svc := &mockService{err: errors.New("db down")}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodPost, "/createAccount", strings.NewReader(`{"email":"user@example.com","password":"secret","role":"tourist"}`))
	w := httptest.NewRecorder()

	rt.CreateAccount(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusInternalServerError)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if !strings.Contains(string(body), "Error: failed to create account.") {
		t.Fatalf("unexpected response body: %s", string(body))
	}
}

func TestRouter_CreateAccount_BadRequest(t *testing.T) {
	svc := &mockService{}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodPost, "/createAccount", strings.NewReader(`not json`))
	w := httptest.NewRecorder()

	rt.CreateAccount(w, req)

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
