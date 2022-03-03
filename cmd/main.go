package main

import (
	"fmt"
	"instalment_plan/api"
	"instalment_plan/config"
	"instalment_plan/sms"
	"instalment_plan/storage"
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/twilio/twilio-go"
)

func main() {
	cfg := config.LoadCfg()

	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPass,
		cfg.PostgresDatabase)

	connDb,err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		panic(err)
	}
	storagePool := storage.NewStorage(connDb)

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username:   cfg.TwilioSid,
		Password:   cfg.TwilioToken,
		AccountSid: cfg.TwilioSid,
	})

	smsClient := sms.NewSmsSerivce(*client)

	server := api.New(api.Option{
		DB:  storagePool.Storage(),
		Sms: *smsClient,
		Conf: cfg,
	})

	err = server.Run(cfg.Port)
	if err != nil {
		panic(err)
	}
}
