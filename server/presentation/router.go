package presentation

import (
	"encoding/json"
	"net/http"
	"server/logic"

	"github.com/gorilla/mux"
)

type GeneralResponse struct {
	Data     interface{}
	Message  interface{}
	Metadata interface{}
}

type Router struct {
	*mux.Router
	db logic.Database
}

func NewRouter(dbClient logic.Database) *Router {
	rt := &Router{
		Router: mux.NewRouter(),
		db:     dbClient,
	}

	rt.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	rt.HandleFunc("/createPlace", rt.CreatePlace).Methods("POST")
	rt.HandleFunc("/getPlaces", rt.GetPlaces).Methods("GET")

	return rt
}

func writeResponse(w http.ResponseWriter, statusCode int, response GeneralResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.WriteHeader(statusCode)
	b, _ := json.Marshal(response)
	w.Write(b)
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	response := GeneralResponse{Data: "", Message: message}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.WriteHeader(statusCode)
	b, _ := json.Marshal(response)
	w.Write(b)
}
