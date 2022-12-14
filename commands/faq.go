package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type faqButton struct {
	display string
	file    string
}

var basePath = "files/faq/"

var faqMenu = [3]faqButton{
	{
		display: "Обо мне",
		file:    basePath + "aboutMe.txt",
	},
	{
		display: "Проф. опыт",
		file:    basePath + "profExp.txt",
	},
	{
		display: "О проектах",
		file:    basePath + "aboutMyProjects.txt",
	},
}

// tests!
func createFAQMenu(split int) tgbotapi.InlineKeyboardMarkup {
	fMenu := make([]tgbotapi.InlineKeyboardButton, len(faqMenu))
	for index, btn := range faqMenu {
		fMenu[index] = tgbotapi.NewInlineKeyboardButtonData(btn.display, btn.file)
	}
	replyKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup()
	keyboard := replyKeyboardMarkup.InlineKeyboard
	threshold := len(fMenu) / split
	for i := 0; i <= threshold; i++ {
		down := i * split
		up := down + split
		if up > cap(fMenu) {
			up = cap(fMenu)
		}
		newSlice := fMenu[down:up]
		keyboard = append(keyboard, newSlice)
	}
	replyKeyboardMarkup.InlineKeyboard = keyboard
	return replyKeyboardMarkup
}
