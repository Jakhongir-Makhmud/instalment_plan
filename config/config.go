package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost string
	PostgresPort int
	PostgresUser string
	PostgresPass string
	PostgresDatabase string
	TwilioSid string
	TwilioToken string
	TwilioPass string
	TwilioUsername string
	Port string
	
}

func LoadCfg() Config {

	var cfg = Config{}
	cfg.Port = cast.ToString(getOrReturnDefault("PORT",":8080"))
	cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST","localhost"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT",5432))
	cfg.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER","jakhongir"))
	cfg.PostgresPass = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD","1"))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE","installment"))
	cfg.TwilioSid = cast.ToString(getOrReturnDefault("TWILIO_SID", "AC21bee4f82b8b3a7f5cb6c0c0d5df1723"))
	cfg.TwilioToken = cast.ToString(getOrReturnDefault("TWILIO_TOKEN", "4a2a7ca4c44d423e69a255d951b1846d"))
	cfg.TwilioUsername = cast.ToString(getOrReturnDefault("TWILIO_USERNAME", "anorboev.jahongir8007@gmail.com"))
	return cfg
}

func getOrReturnDefault(env_var string, def interface{}) interface{} {
	val, exists := os.LookupEnv(env_var)
	if exists {
		return val
	}
	return def
}