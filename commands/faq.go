package commands

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

func createFAQMenu(split int, menu []faqButton) (tgbotapi.InlineKeyboardMarkup, error) {
	if split < 0 {
		return tgbotapi.NewInlineKeyboardMarkup(), errors.New("value is below zero")
	}
	fMenu := make([]tgbotapi.InlineKeyboardButton, len(menu))
	for index, btn := range menu {
		fMenu[index] = tgbotapi.NewInlineKeyboardButtonData(btn.display, btn.file)
	}
	replyKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup()
	threshold := len(fMenu) / split
	var keyboard [][]tgbotapi.InlineKeyboardButton
	for i := 0; i < threshold; i++ {
		down := i * split
		up := down + split
		if up > len(fMenu) {
			up = len(fMenu)
		}
		keyboard = append(keyboard, fMenu[down:up])
	}
	if remainder := len(fMenu) - split*threshold; remainder > 0 {
		keyboard = append(keyboard, fMenu[len(fMenu)-remainder:])
	}
	replyKeyboardMarkup.InlineKeyboard = keyboard
	return replyKeyboardMarkup, nil
}
