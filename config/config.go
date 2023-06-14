package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	ChargeApiUrl string
}

var config *Config

func Load() *Config {
	err := godotenv.Load("dev.env")
	if err == nil {
		log.Infoln("Load dev.env file for local dev")
	}
	if config == nil {
		if os.Getenv("CHARGE_API_URL") == "" { //other env value might not set as well
			_ = fmt.Errorf("CHARGE_API_URL is not set:")
		}

		config = &Config{
			ChargeApiUrl: os.Getenv("CHARGE_API_URL"),
		}
	}
	return config
}
