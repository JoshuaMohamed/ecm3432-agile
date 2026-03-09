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
	service logic.Service
}

func NewRouter(service logic.Service) *Router {
	rt := &Router{
		Router:  mux.NewRouter(),
		service: service,
	}

	rt.Use(corsMiddleware)

	rt.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	rt.HandleFunc("/createPlace", rt.CreatePlace).Methods("POST", "OPTIONS")
	rt.HandleFunc("/getPlaces", rt.GetPlaces).Methods("GET", "OPTIONS")
	rt.HandleFunc("/signup", rt.SignUp).Methods("POST", "OPTIONS")
	rt.HandleFunc("/login", rt.LogIn).Methods("POST", "OPTIONS")

	return rt
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func writeResponse(w http.ResponseWriter, statusCode int, response GeneralResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	b, _ := json.Marshal(response)
	w.Write(b)
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	response := GeneralResponse{Data: "", Message: message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	b, _ := json.Marshal(response)
	w.Write(b)
}
