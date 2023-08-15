package main

import (
	"log"

	"github.com/AnNosov/tele_bot/config"
	"github.com/AnNosov/tele_bot/enternal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
