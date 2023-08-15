package app

import (
	"log"

	"github.com/AnNosov/tele_bot/config"
	"github.com/AnNosov/tele_bot/enternal/controller/bot"
	"github.com/AnNosov/tele_bot/enternal/usecase"
	"github.com/AnNosov/tele_bot/enternal/usecase/repo"
	"github.com/AnNosov/tele_bot/pkg/postgres"
	"github.com/AnNosov/tele_bot/pkg/tlgrm"
)

func Run(cfg *config.Config) {

	postgres, err := postgres.New(&cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	pgRepo := repo.New(postgres)
	teleBot, err := tlgrm.New(&cfg.Telebot)
	if err != nil {
		log.Fatal(err)
	}

	gameAction := usecase.New(pgRepo, teleBot)

	bot.BotController(*gameAction)

}
