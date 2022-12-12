package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

const botKeyApi = "ABOUT_ME_BOT_API_KEY"

func main() {
	bot := makeBot(true)

	updates := makeUpdateChan(bot)

	for update := range updates {

		if update.Message == nil {
			continue
		}

		go makeReply(update, bot)
	}

}

func makeReply(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
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
