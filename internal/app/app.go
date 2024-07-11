package app

import (
	"log"
	"net/http"
	"os"

	"todo-list/internal/database"
	"todo-list/internal/handlers"
)

func Run() {
	database.InitDB()

	deps := handler.Dependencies{
		DB: database.DB,
	}

	h, err := handler.New(deps, handler.WithHTTPHandler())
	if err != nil {
		log.Fatalf("Error initializing handler: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, h.HTTP); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
