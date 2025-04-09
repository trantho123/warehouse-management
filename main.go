package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/trantho123/warehouse-management/utils"
)

func ConnectDB(config *utils.Config) (*sql.DB, error) {
	// Open database connection
	db, err := sql.Open("postgres", config.DBSource)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	return db, nil
}

func main() {
	fmt.Println("Welcome to the Warehouse Management System!")

	// Create database configuration
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// Connect to database
	db, err := ConnectDB(&config)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	fmt.Println("Successfully connected to database!")

}
