package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"todo-list/util"

	_ "github.com/lib/pq"

	"github.com/pressly/goose/v3"
)

var DB *sql.DB

func InitDB() {
	var err error

	// Load configuration
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Print DB_SOURCE for debugging
	fmt.Println("DB_SOURCE:", config.DBSource)

	// Open a connection to the database using DB_SOURCE directly
	DB, err = sql.Open("postgres", config.DBSource)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Check if the connection to the database is working
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	// Run database migrations
	if err := runMigrations(); err != nil {
		log.Fatalf("Could not run database migrations: %v", err)
	}

	log.Println("Connected to the database successfully!")
}

func runMigrations() error {
	// Get the absolute path for the migrations directory
	dir, err := filepath.Abs("./db/migrations")
	if err != nil {
		return fmt.Errorf("could not get absolute path for migrations: %w", err)
	}

	// Check if the migrations directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Printf("Migrations directory '%s' does not exist, skipping migrations", dir)
		return nil // No migrations to run
	}

	// Print the contents of the migrations directory
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("could not read migrations directory: %w", err)
	}

	log.Println("Migration files:")
	for _, file := range files {
		log.Println(file.Name())
	}

	// Perform database migrations
	if err := goose.Up(DB, dir); err != nil {
		return fmt.Errorf("error running migrations: %w", err)
	}

	log.Println("Database migrations ran successfully!")
	return nil
}
