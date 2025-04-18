package repository

import (
  "database/sql"
  "fmt"
  "log"
  "os"

  _ "github.com/lib/pq"
)

func InitDatabase() (*sql.DB, error) {
  env := os.Getenv("ENV")
  prod := os.Getenv("PROD_DB_URL")
  dev := os.Getenv("DEV_DB_URL")

  var connStr string

  switch env {
  case "production":
	  if prod == "" {
		  log.Fatal("Missing PROD_DB_URL in production")
	  }
	  connStr = prod
  default:
	  if dev == "" {
		  log.Fatal("Missing DEV_DB_URL in development")
	  }
	  connStr = dev
  }

  if connStr == "" {
		return nil, fmt.Errorf("missing database URL for environment: %s", env)
	}

  db, err := sql.Open("postgres", connStr)
  if err != nil {
    return nil, fmt.Errorf("failed to connect to database: %v", err)
  }
  log.Println("Database connection established.")
  if err := db.Ping(); err != nil {
    return nil, fmt.Errorf("failed to ping database: %v", err)
  }
  return db, nil
}

func CloseDatabase(db *sql.DB) {
  if db != nil {
    if err := db.Close(); err != nil {
      log.Printf("Error closing the database: %v", err)
    } else {
      log.Printf("Database connection closed.")
    }
  }
}