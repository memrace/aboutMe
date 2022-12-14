package main

import (
	"aboutMe/commands"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

const botKeyApi = "ABOUT_ME_BOT_API_KEY"

func main() {
	bot := makeBot(true)

	updates := makeUpdateChan(bot)

	for update := range updates {

		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		if update.CallbackQuery != nil {
			go makeReplyForCallback(&update, bot)
			continue
		}

		if update.Message != nil {
			go makeReply(&update, bot)
			continue
		}

	}

}
func makeReplyForCallback(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	commands.MakeCommandHandler(bot, update).ProcessCallback()
}
func makeReply(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	commands.MakeCommandHandler(bot, update).Process()
}

func makeUpdateChan(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)
	return updates
}

func makeBot(debug bool) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(os.Getenv(botKeyApi))
	if err != nil {
		panic(err)
	}

	bot.Debug = debug
	return bot
}
