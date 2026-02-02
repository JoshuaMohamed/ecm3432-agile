package main

import (
	"log"
	"log/slog"
	"net/http"
	"server/logic"
	"server/persistence"
	"server/presentation"
)

const PORT = "8080"

func main() {
	// 1. Connect to database (create if missing)
	dbClient, err := persistence.NewDatabaseClient()
	if err != nil {
		log.Fatal(err)
	}
	defer dbClient.Close()

	// 2. Create Places table if missing
	err = dbClient.CreateTable(logic.TableDetails{
		Name: "Places",
		Schema: `
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			postcode TEXT NOT NULL
		`,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 3. Create router and pass database client
	router := presentation.NewRouter(dbClient)

	// 4. Start server
	slog.Info("Server is starting,", "port", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
