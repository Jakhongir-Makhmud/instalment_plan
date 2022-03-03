package handlers

import (
	"instalment_plan/config"
	"instalment_plan/sms"
	"instalment_plan/storage/repo"
)

type handlers struct {
	cfg config.Config
	db repo.DatabaseRepo
	sms sms.Sms
}

func NewHandler(cfg config.Config,db repo.DatabaseRepo,sms sms.Sms) *handlers {
	return &handlers{
		cfg: cfg,
		db: db,
		sms: sms,
	}
}
