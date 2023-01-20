package main

import (
	service "aboutMe/api"
	"aboutMe/commands"
	"context"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const botKeyApi = "ABOUT_ME_BOT_API_KEY"

func main() {

	service := service.MakeService()

	bot := makeBot(true)
	updates := makeUpdateChan(bot)
	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}
		if update.CallbackQuery != nil {
			go commands.MakeCommandHandler(&service, bot, &update).ProcessCallback()
			continue
		}
		if update.Message != nil {
			go commands.MakeCommandHandler(&service, bot, &update).Process(context.Background())
			continue
		}
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
