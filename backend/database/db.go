package database

import (
        "database/sql"
        "fmt"
        "log"

        // Import the necessary database driver (e.g., PostgreSQL, MySQL, SQLite, etc.)
        _ "github.com/lib/pq"
)

var db *sql.DB

// InitDB initializes the database connection.
func InitDB(dataSourceName string) error {
        var err error
        db, err = sql.Open("postgres", dataSourceName)
        if err != nil {
                return fmt.Errorf("failed to connect to database: %v", err)
        }

        if err = db.Ping(); err != nil {
                return fmt.Errorf("failed to ping database: %v", err)
        }

        log.Println("Connected to the database!")
        return nil
}

// CloseDB closes the database connection.
func CloseDB() {
        if db != nil {
                db.Close()
                log.Println("Database connection closed.")
        }
}
