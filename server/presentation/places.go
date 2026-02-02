package presentation

import (
	"fmt"
	"log/slog"
	"net/http"
	"server/logic"
)

// GetPlaces handles GET requests for places
func (rt *Router) GetPlaces(w http.ResponseWriter, req *http.Request) {
	postcode := req.URL.Query().Get("postcode")
	filter := req.URL.Query().Get("filter")

	data, err := logic.GetPlaces(rt.db, postcode, filter, 100, 0)
	if err != nil {
		slog.Error("Failed to get places.", "error", err)
		writeErrorResponse(w, http.StatusInternalServerError, "Error: failed to get places. Check the postcode and filter.")
		return
	}

	message := fmt.Sprintf("Got places for postcode %s", postcode)
	slog.Info(message)
	writeResponse(w, http.StatusOK, GeneralResponse{
		Data:    data,
		Message: message,
	})
}
