package presentation

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"server/logic"
	"time"
)

// SignUp handles POST requests for creating a new account
func (rt *Router) SignUp(w http.ResponseWriter, req *http.Request) {
	var account logic.Account

	err := json.NewDecoder(req.Body).Decode(&account)
	if err != nil {
		slog.Error("Failed to decode request body", "error", err)
		writeErrorResponse(w, http.StatusBadRequest, "Error: Bad Request")
		return
	}

	session, err := rt.service.SignUp(account)
	if err != nil {
		slog.Error("Failed to create account", "error", err)
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session.Token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	message := "Created account"
	slog.Info(message)
	writeResponse(w, http.StatusOK, GeneralResponse{
		Data:    "",
		Message: message,
	})
}

// LogIn handles POST requests for assuming an existing account
func (rt *Router) LogIn(w http.ResponseWriter, req *http.Request) {
	var account logic.Account

	err := json.NewDecoder(req.Body).Decode(&account)
	if err != nil {
		slog.Error("Failed to decode request body", "error", err)
		writeErrorResponse(w, http.StatusBadRequest, "Error: Bad Request")
		return
	}

	session, err := rt.service.LogIn(account)
	if err != nil {
		slog.Error("Failed to authenticate credentials", "error", err)
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session.Token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	message := "Successfully logged in"
	slog.Info(message)
	writeResponse(w, http.StatusOK, GeneralResponse{
		Data:    "",
		Message: message,
	})
}

// LogOut handles DELETE requests for deleting session
func (rt *Router) LogOut(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session_token")
	if err != nil || cookie.Value == "" {
		writeErrorResponse(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	err = rt.service.LogOut(cookie.Value)
	if err != nil {
		slog.Error("Failed to validate session", "error", err)
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Expire the session token cookie on successful session deletion.
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})

	message := "Successfully logged out"
	slog.Info(message)
	writeResponse(w, http.StatusOK, GeneralResponse{
		Data:    "",
		Message: message,
	})
}

// ValidateSession returns the email for the current session token.
func (rt *Router) ValidateSession(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session_token")
	if err != nil || cookie.Value == "" {
		writeErrorResponse(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	email, err := rt.service.ValidateSession(cookie.Value)
	if err != nil {
		writeErrorResponse(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	writeResponse(w, http.StatusOK, GeneralResponse{
		Data: map[string]string{"email": email},
		Message: "Session valid",
	})
}
