package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"
	"server/logic"
	"server/persistence"
	"server/presentation"
)

const PORT = "8080"

func main() {
	// Connect to database (create if missing)
	dbClient, err := persistence.NewDatabaseClient()
	if err != nil {
		log.Fatal(err)
	}
	defer dbClient.Close()

	// Load table schemas from config
	schemaFile, err := os.ReadFile("schema.json")
	if err != nil {
		log.Fatal(err)
	}

	var tables []logic.TableDetails
	if err := json.Unmarshal(schemaFile, &tables); err != nil {
		log.Fatal(err)
	}

	// Create tables
	for _, table := range tables {
		if err := dbClient.CreateTable(table); err != nil {
			log.Fatal(err)
		}
	}

	// Create service and router
	service := &logic.ServiceImpl{DB: dbClient}
	router := presentation.NewRouter(service)

	// Start server
	slog.Info("Server is starting,", "port", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
