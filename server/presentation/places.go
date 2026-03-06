package presentation

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"server/logic"
)


// CreatePlace handles POST requests for adding a new place
func (rt *Router) CreatePlace(w http.ResponseWriter, req *http.Request) {
	var place logic.Place

	err := json.NewDecoder(req.Body).Decode(&place)
	if err != nil {
		slog.Error("Failed to decode request body.", "error", err)
		writeErrorResponse(w, http.StatusBadRequest, "Error: Bad Request.")
		return
	}

	err = rt.service.CreatePlace(place)
	if err != nil {
		slog.Error("Failed to create place.", "error", err)
		writeErrorResponse(w, http.StatusInternalServerError, "Error: failed to create place.")
		return
	}

	message := "Created place."
	slog.Info(message)
	writeResponse(w, http.StatusOK, GeneralResponse{
		Data:    "",
		Message: message,
	})
}

// GetPlaces handles GET requests for places
func (rt *Router) GetPlaces(w http.ResponseWriter, req *http.Request) {
	postcode := req.URL.Query().Get("postcode")
	filter := req.URL.Query().Get("filter")

	data, err := rt.service.GetPlaces(postcode, filter, 100, 0)
	if err != nil {
		slog.Error("Failed to get places.", "error", err)
		writeErrorResponse(w, http.StatusInternalServerError, "Error: Failed to get places. Check the postcode and filter.")
		return
	}

	message := fmt.Sprintf("Got places for postcode %s", postcode)
	slog.Info(message)
	writeResponse(w, http.StatusOK, GeneralResponse{
		Data:    data,
		Message: message,
	})
}
