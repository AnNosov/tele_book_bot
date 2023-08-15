package tlgrm

import (
	"github.com/AnNosov/tele_bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TlgBot struct {
	TB *tgbotapi.BotAPI
}

func New(cfg *config.Telebot) (*TlgBot, error) {

	tBot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}

	tb := &TlgBot{
		TB: tBot,
	}

	return tb, nil

}
