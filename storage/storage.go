package storage

import (
	"instalment_plan/storage/database"
	"instalment_plan/storage/repo"

	"github.com/jmoiron/sqlx"
)

// Storage is only used to store,given informations
// There is no calculations or logic in database methods
type storagePool struct {
	db       *sqlx.DB
	database repo.DatabaseRepo
}

// NewStorage is used to initialize database
func NewStorage(db *sqlx.DB) *storagePool {
	return &storagePool{
		db:       db,
		database: database.NewDatabase(db),
	}
}

// Storage returns DatabaseRepo interface, from which you can use db methods
func (s storagePool) Storage() repo.DatabaseRepo {
	return s.database
}
