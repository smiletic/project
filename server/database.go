package server

import (
	"database/sql"
	"fmt"
	"masterRad/config"
	"masterRad/db"

	_ "github.com/lib/pq"
)

// InitializeDb initializes connection to database.
func InitializeDb() error {
	dbHandle, err := sql.Open("postgres", config.GetDatabaseConnectionString())
	if err != nil {
		return err
	}

	// Initialize connection pool
	dbHandle.SetMaxIdleConns(config.GetDatabaseMaxIdleConnections())
	dbHandle.SetMaxOpenConns(config.GetDatabaseMaxOpenConnections())
	dbHandle.SetConnMaxLifetime(config.GetDatabaseConnectionMaxLifetime())

	fmt.Printf("\tDatabase connection pool: SetMaxIdleConns = %v\n", config.GetDatabaseMaxIdleConnections())
	fmt.Printf("\tDatabase connection pool: SetMaxOpenConns = %v\n", config.GetDatabaseMaxOpenConnections())
	fmt.Printf("\tDatabase connection pool: SetConnMaxLifetime = %v\n", config.GetDatabaseConnectionMaxLifetime())

	err = dbHandle.Ping()
	if err != nil {
		return err
	}

	// Assign global database handle
	db.Handle = dbHandle

	return nil
}
