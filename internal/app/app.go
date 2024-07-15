package app

import (
	"log"
	"net/http"
	"os"
	"time"
	"todo-list/internal/database"
	myHttp "todo-list/internal/handlers/http"
	"github.com/go-chi/chi/v5"
	"github.com/hellofresh/health-go/v5"
	healthPg "github.com/hellofresh/health-go/v5/checks/postgres"
)

func Run() {
	// Initialize the database
	database.InitDB()

	// Create a new chi router
	router := chi.NewRouter()

	// Register task routes
	myHttp.RegisterTaskRoutes(router) // Correct usage

	// Set up health checks
	healthHandler, _ := health.New(health.WithComponent(health.Component{
		Name:    "todo-list",
		Version: "v1.0",
	}), health.WithChecks(
		health.Config{
			Name:      "postgres",
			Timeout:   time.Second * 2,
			SkipOnErr: false,
			Check: healthPg.New(healthPg.Config{
				DSN: os.Getenv("DB_SOURCE"),
			}),
		},
	))

	// Register health check endpoint
	router.Get("/status", healthHandler.HandlerFunc)

	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
