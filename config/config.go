package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost string
	PostgresPort string
	TwilioSid string
	TwilioToken string
	TwilioPass string
	TwilioUsername string
}

func LoadCfg() Config {

	var cfg = Config{}

	cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST","localhost"))
	cfg.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT","5432"))
	cfg.TwilioSid = cast.ToString(getOrReturnDefault("TWILIO_SID", "AC21bee4f82b8b3a7f5cb6c0c0d5df1723"))
	cfg.TwilioToken = cast.ToString(getOrReturnDefault("TWILIO_TOKEN", "c201c94636e5c3387f50502d9bfd92e7"))
	cfg.TwilioUsername = cast.ToString(getOrReturnDefault("TWILIO_USERNAME", "anorboev.jahongir8007@gmail.com"))
	return cfg
}

func getOrReturnDefault(env_var,def string) interface{} {
	val, exists := os.LookupEnv(env_var)
	if exists {
		return val
	}
	return def
}