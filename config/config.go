package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func Load() {
	err := godotenv.Load("dev.env")
	if err == nil {
		log.Infoln("Load dev.env file for local dev")
	}
	return
}
