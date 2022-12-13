package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const readMe = "readme"

const start = "start"

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
}

type commandHandler struct {
	bot    *tgbotapi.BotAPI
	update *tgbotapi.Update
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
		go showReadMe(message.Chat.ID, bot)
		return
	case start:
		go showWelcome(message.Chat.ID, bot)
		return
	}
}

func workWithSimpleMessage(chatID int64, bot *tgbotapi.BotAPI) {
	sendMessage(bot, makeMessage(chatID, "¯\\_(ツ)_/¯"))
}

func showReadMe(chatID int64, bot *tgbotapi.BotAPI) {
	sendMessage(bot, makeMessage(chatID, getTextFromReadMe()))
}

func showWelcome(chatID int64, bot *tgbotapi.BotAPI) {
	sendMessage(bot, makeMessage(chatID, "Приветствую! \nПредлагаю вызвать команду readme для первичного ознакомления ;)"))
}
