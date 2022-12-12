package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func sendMessage(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
func makeMessage(chatId int64, text string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatId, text)
}
