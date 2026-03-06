package presentation

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"server/logic"
)

// GetAccounts handles POST requests for creating a new account
func (rt *Router) CreateAccount(w http.ResponseWriter, req *http.Request) {
	var account logic.Account

	err := json.NewDecoder(req.Body).Decode(&account)
	if err != nil {
		slog.Error("Failed to decode request body.", "error", err)
		writeErrorResponse(w, http.StatusBadRequest, "Error: Bad Request.")
		return
	}

	err = rt.service.CreateAccount(account)
	if err != nil {
		slog.Error("Failed to create account.", "error", err)
		writeErrorResponse(w, http.StatusConflict, err.Error())
		return
	}

	message := "Created account."
	slog.Info(message)
	writeResponse(w, http.StatusOK, GeneralResponse{
		Data:    "",
		Message: message,
	})
}
