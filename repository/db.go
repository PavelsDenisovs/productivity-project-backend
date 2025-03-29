package repository

import (
  "database/sql"
  "fmt"
  "log"
  "os"

  _ "github.com/lib/pq"
)

func InitDatabase() (*sql.DB, error) {
  connStr := os.Getenv("DB_URL")
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