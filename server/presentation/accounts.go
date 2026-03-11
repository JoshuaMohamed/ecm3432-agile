package presentation

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"server/logic"
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

	err = rt.service.LogIn(account)
	if err != nil {
		slog.Error("Failed to authenticate credentials", "error", err)
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	message := "Successfully logged in"
	slog.Info(message)
	writeResponse(w, http.StatusOK, GeneralResponse{
		Data:    "",
		Message: message,
	})
}
