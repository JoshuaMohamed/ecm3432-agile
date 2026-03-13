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

func TestRouter_SignUp_Success(t *testing.T) {
	svc := &mockService{}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(`{"email":"user@example.com","password":"secret","role":"tourist"}`))
	w := httptest.NewRecorder()

	rt.SignUp(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusOK)
	}
}

func TestRouter_SignUp_DBError(t *testing.T) {
	svc := &mockService{err: errors.New("db down")}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(`{"email":"user@example.com","password":"secret","role":"tourist"}`))
	w := httptest.NewRecorder()

	rt.SignUp(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusInternalServerError)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if !strings.Contains(string(body), "db down") {
		t.Fatalf("unexpected response body: %s", string(body))
	}
}

func TestRouter_SignUp_BadRequest(t *testing.T) {
	svc := &mockService{}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(`not json`))
	w := httptest.NewRecorder()

	rt.SignUp(w, req)

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

func TestRouter_LogIn_Success(t *testing.T) {
	svc := &mockService{}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email":"user@example.com","password":"secret"}`))
	w := httptest.NewRecorder()

	rt.LogIn(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusOK)
	}
}

func TestRouter_LogIn_DBError(t *testing.T) {
	svc := &mockService{err: errors.New("db down")}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email":"user@example.com","password":"secret"}`))
	w := httptest.NewRecorder()

	rt.LogIn(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusInternalServerError)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if !strings.Contains(string(body), "db down") {
		t.Fatalf("unexpected response body: %s", string(body))
	}
}

func TestRouter_LogIn_BadRequest(t *testing.T) {
	svc := &mockService{}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`not json`))
	w := httptest.NewRecorder()

	rt.LogIn(w, req)

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

func TestRouter_LogOut_Success(t *testing.T) {
	svc := &mockService{}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodDelete, "/logout", nil)
	req.AddCookie(&http.Cookie{Name: "session_token", Value: "abc123"})
	w := httptest.NewRecorder()

	rt.LogOut(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusOK)
	}

	foundExpiredCookie := false
	for _, c := range res.Cookies() {
		if c.Name == "session_token" && c.MaxAge < 0 {
			foundExpiredCookie = true
			break
		}
	}
	if !foundExpiredCookie {
		t.Fatalf("expected expired session_token cookie")
	}
}

func TestRouter_LogOut_DBError(t *testing.T) {
	svc := &mockService{err: errors.New("db down")}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodDelete, "/logout", nil)
	req.AddCookie(&http.Cookie{Name: "session_token", Value: "abc123"})
	w := httptest.NewRecorder()

	rt.LogOut(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusInternalServerError)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if !strings.Contains(string(body), "db down") {
		t.Fatalf("unexpected response body: %s", string(body))
	}
}

func TestRouter_LogOut_Unauthorized_NoCookie(t *testing.T) {
	svc := &mockService{}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodDelete, "/logout", nil)
	w := httptest.NewRecorder()

	rt.LogOut(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusUnauthorized)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if !strings.Contains(string(body), "Invalid session") {
		t.Fatalf("unexpected response body: %s", string(body))
	}
}

func TestRouter_ValidateSession_Success(t *testing.T) {
	svc := &mockService{email: "user@example.com"}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/session", nil)
	req.AddCookie(&http.Cookie{Name: "session_token", Value: "abc123"})
	w := httptest.NewRecorder()

	rt.ValidateSession(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusOK)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if !strings.Contains(string(body), "user@example.com") {
		t.Fatalf("unexpected response body: %s", string(body))
	}
}

func TestRouter_ValidateSession_Unauthorized_NoCookie(t *testing.T) {
	svc := &mockService{}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/session", nil)
	w := httptest.NewRecorder()

	rt.ValidateSession(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusUnauthorized)
	}
}

func TestRouter_ValidateSession_Unauthorized_ServiceError(t *testing.T) {
	svc := &mockService{err: errors.New("invalid")}
	rt := presentation.NewRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/session", nil)
	req.AddCookie(&http.Cookie{Name: "session_token", Value: "abc123"})
	w := httptest.NewRecorder()

	rt.ValidateSession(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusUnauthorized)
	}
}
