package commands

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const readMe = "readme"

const start = "start"

const faq = "faq"

func MakeCommandHandler(
	bot *tgbotapi.BotAPI,
	update *tgbotapi.Update,
) CommandHandler {
	return &commandHandler{
		bot:    bot,
		update: update,
	}
}

type CommandHandler interface {
	Process()
	ProcessCallback()
}

type commandHandler struct {
	bot    *tgbotapi.BotAPI
	update *tgbotapi.Update
}

func (handler *commandHandler) ProcessCallback() {
	callbackQuery := handler.update.CallbackQuery
	var btn *faqButton
	for _, b := range faqMenu {
		if b.command == callbackQuery.Data {
			btn = &b
			break
		}
	}
	if btn == nil {
		panic(errors.New("There is not any such button that has such data " + callbackQuery.Data))
	}
	callbackRequest := tgbotapi.NewCallback(callbackQuery.ID, btn.display)
	response, err := handler.bot.Request(callbackRequest)
	if err != nil {
		println(err)
	}
	if response.Ok {
		sendMessage(handler.bot, makeMessage(callbackQuery.Message.Chat.ID, getTextFromFile(callbackQuery.Data)))
	}

}

func (handler *commandHandler) Process() {
	message := handler.update.Message
	if message.IsCommand() {
		workWithCommand(handler.bot, handler.update)
	} else {
		workWithSimpleMessage(message.Chat.ID, handler.bot)
	}
}

func workWithCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	message := update.Message
	command := message.Command()
	switch command {
	case readMe:
		go sendMessage(bot, makeMessage(message.Chat.ID, getTextFromFile("README.md")))
		return
	case start:
		go sendMessage(bot, makeMessage(message.Chat.ID, "Приветствую! \nПредлагаю вызвать команду readme для первичного ознакомления ;)"))
		return
	case faq:
		go func() {
			message := tgbotapi.NewMessage(message.Chat.ID, message.Text)
			if menu, err := createFAQMenu(3, faqMenu[:]); err != nil {
				println(err)
			} else {
				message.ReplyMarkup = menu
				sendMessage(bot, message)
			}
		}()
		return
	}
}

func workWithSimpleMessage(chatID int64, bot *tgbotapi.BotAPI) {
	sendMessage(bot, makeMessage(chatID, "¯\\_(ツ)_/¯"))
}
