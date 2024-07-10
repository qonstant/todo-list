package app

import (
	"log"
	"net/http"
	"os"

	"todo-list/internal/database"
	"todo-list/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func Run() {
	database.InitDB()

	router := chi.NewRouter()
	handlers.RegisterTaskRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
