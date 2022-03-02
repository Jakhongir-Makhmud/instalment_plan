package storage

import (
	"instalment_plan/storage/database"
	"instalment_plan/storage/repo"

	"github.com/jmoiron/sqlx"
)

type storagePool struct {
	db       *sqlx.DB
	database repo.DatabaseRepo
}

func NewStorage(db *sqlx.DB) *storagePool {
	return &storagePool{
		db:       db,
		database: database.NewDatabase(db),
	}
}

func (s storagePool) Storage() repo.DatabaseRepo {
	return s.database
}
