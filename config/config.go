package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	ChatGPTToken string
}

var config *Config

func Load() *Config {
	err := godotenv.Load("dev.env")
	if err == nil {
		log.Infoln("Load dev.env file for local dev")
	}

	if config == nil {
		if os.Getenv("CHATGPT_TOKEN") == "" { //other env value might not set as well
			log.Fatalln("CHATGPT_TOKEN is not set:")
		}

		config = &Config{
			ChatGPTToken: os.Getenv("CHATGPT_TOKEN"),
		}
	}
	fmt.Printf("LOAD CONFIG:%v", config)
	return config
}
