package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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
	// Directory where migration files are stored
	dir := "./db/migrations"

	// Check if the migrations directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil // No migrations to run
	}

	// Set the base file system for migrations
	goose.SetBaseFS(os.DirFS(dir))

	// Perform database migrations
	if err := goose.Up(DB, ""); err != nil {
		return err
	}

	return nil
}
