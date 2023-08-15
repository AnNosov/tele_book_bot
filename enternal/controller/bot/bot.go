package bot

import (
	"log"
	"strconv"
	"strings"

	"github.com/AnNosov/tele_bot/enternal/usecase"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup()
var schemaPrefix = "schema:"
var elementPrefix = "element:"

func deleteSchemaPrefix(s string) string {
	str := strings.TrimPrefix(s, schemaPrefix)
	return str
}

func deleteElementPrefix(s string) string {
	str := strings.TrimPrefix(s, elementPrefix)
	return str
}

func setSchemaPrefix(s string) string {
	return schemaPrefix + s
}

func setElementPrefix(s string) string {
	return elementPrefix + s
}

func BotController(gA usecase.GameAction) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := gA.Tlgrm.TB.GetUpdatesChan(u)

	go func() {

		for update := range updates {
			if update.CallbackQuery != nil {
				buttonValue := update.CallbackQuery.Data
				inlineKeyboard.InlineKeyboard = [][]tgbotapi.InlineKeyboardButton{} // clear keyboard

				if strings.HasPrefix(buttonValue, schemaPrefix) {
					schema, err := strconv.Atoi(deleteSchemaPrefix(buttonValue))

					if err != nil {
						log.Panic(err)
					}
					element, err := gA.FirstStepElement(schema)

					if err != nil {
						log.Panic(err)
					}

					message := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, element.Desc)
					//message.ParseMode = "HTML"

					for id, info := range element.Next {
						btn := tgbotapi.NewInlineKeyboardButtonData(info, setElementPrefix(strconv.Itoa(id)))
						inlineKeyboard.InlineKeyboard = append(inlineKeyboard.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(btn)) //append(inlineKeyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{btn})
					}

					message.ReplyMarkup = inlineKeyboard
					_, err = gA.Tlgrm.TB.Send(message)
					if err != nil {
						log.Panic(err)
					}

				} else if strings.HasPrefix(buttonValue, elementPrefix) {
					elementId, err := strconv.Atoi(deleteElementPrefix(buttonValue))
					if err != nil {
						log.Panic(err)
					}
					element, err := gA.StepInfo(elementId)
					if err != nil {
						log.Panic(err)
					}

					if len(element.Next) != 0 {

						message := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, element.Desc)
						//message.ParseMode = "markdown"

						for id, info := range element.Next {
							btn := tgbotapi.NewInlineKeyboardButtonData(info, setElementPrefix(strconv.Itoa(id)))
							inlineKeyboard.InlineKeyboard = append(inlineKeyboard.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(btn))
						}

						message.ReplyMarkup = inlineKeyboard
						_, err = gA.Tlgrm.TB.Send(message)
						if err != nil {
							log.Panic(err)
						}
					} else {
						books, err := gA.BookList()
						if err != nil {
							log.Panic(err)
						}

						message := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Выберите книгу: ")
						//message.ParseMode = "Markdown"

						for _, book := range books {
							btn := tgbotapi.NewInlineKeyboardButtonData(book.Name, setSchemaPrefix(strconv.Itoa(book.Id)))
							inlineKeyboard.InlineKeyboard = append(inlineKeyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{btn})
						}

						message.ReplyMarkup = inlineKeyboard
						_, err = gA.Tlgrm.TB.Send(message)
						if err != nil {
							log.Panic(err)
						}
					}

				} else {
					log.Println("ERROR при обработке префикса, button value: ", buttonValue)
				}

			} else {
				inlineKeyboard.InlineKeyboard = [][]tgbotapi.InlineKeyboardButton{} // clear keyboard

				books, err := gA.BookList()
				if err != nil {
					log.Panic(err)
				}

				message := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите книгу: ")
				//message.ParseMode = "Markdown"

				for _, book := range books {
					btn := tgbotapi.NewInlineKeyboardButtonData(book.Name, setSchemaPrefix(strconv.Itoa(book.Id)))
					inlineKeyboard.InlineKeyboard = append(inlineKeyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{btn})
				}

				message.ReplyMarkup = inlineKeyboard
				_, err = gA.Tlgrm.TB.Send(message)
				if err != nil {
					log.Panic(err)
				}

			}

		}

	}()
	select {}

}
