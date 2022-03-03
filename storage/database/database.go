package database

import (
	"github.com/jmoiron/sqlx"
)

// Main struct to interact with database
type Database struct {
	db *sqlx.DB
}

// NewDatabase function initalizes main Database struct
func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}
