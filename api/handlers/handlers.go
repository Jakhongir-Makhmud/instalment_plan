package handlers

import (
	"instalment_plan/storage/repo"
)

type handlers struct {
	db repo.DatabaseRepo
}

func NewHandler(db repo.DatabaseRepo) *handlers {
	return &handlers{
		db: db,
	}
}
