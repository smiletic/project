package db

import "database/sql"

// DefaultTx represents default transaction options
var DefaultTx = &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false}

// DefaultReadOnlyTx represents default transaction options for read-only transactions
var DefaultReadOnlyTx = &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: true}

// Handle is the database handle
// TODO: Remove this! This var should be in the main project, not in the library
var Handle *sql.DB
