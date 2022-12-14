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
		if b.file == callbackQuery.Data {
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
		go showReadMe(message.Chat.ID, bot)
		return
	case start:
		go showWelcome(message.Chat.ID, bot)
		return
	case faq:
		go showFaqMenu(message.Chat.ID, bot, message)
		return
	}
}

func workWithSimpleMessage(chatID int64, bot *tgbotapi.BotAPI) {
	sendMessage(bot, makeMessage(chatID, "¯\\_(ツ)_/¯"))
}

func showFaqMenu(chatID int64, bot *tgbotapi.BotAPI, updateMessage *tgbotapi.Message) {
	message := tgbotapi.NewMessage(chatID, updateMessage.Text)
	message.ReplyMarkup = createFAQMenu(3)
	sendMessage(bot, message)
}

func showReadMe(chatID int64, bot *tgbotapi.BotAPI) {
	sendMessage(bot, makeMessage(chatID, getTextFromFile("README.md")))
}

func showWelcome(chatID int64, bot *tgbotapi.BotAPI) {
	sendMessage(bot, makeMessage(chatID, "Приветствую! \nПредлагаю вызвать команду readme для первичного ознакомления ;)"))
}
