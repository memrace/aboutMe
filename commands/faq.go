package commands

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type faqButton struct {
	display string
	command string
}

var basePath = "files/faq/"

var faqMenu = [...]faqButton{
	{
		display: "Обо мне",
		command: basePath + "aboutMe.txt",
	},
	{
		display: "Проф. опыт",
		command: basePath + "profExp.txt",
	},
	{
		display: "О проектах",
		command: basePath + "aboutMyProjects.txt",
	},
	{
		display: "О команде",
		command: basePath + "team.txt"},
	{
		display: "Контакты",
		command: basePath + "contacts.txt",
	},
	{
		display: "Прошлый стек технологий",
		command: basePath + "prevStack.txt",
	},
	{
		display: "С чем есть опыт в Go",
		command: basePath + "goExp.txt",
	},
	{
		display: "Почему Go?",
		command: basePath + "whyGo.txt",
	},
}

func createFAQMenu(split int, menu []faqButton) (tgbotapi.InlineKeyboardMarkup, error) {
	if split < 0 {
		return tgbotapi.NewInlineKeyboardMarkup(), errors.New("split is below a zero")
	}
	if split == 0 {
		return tgbotapi.NewInlineKeyboardMarkup(), errors.New("split is a zero")
	}
	fMenu := make([]tgbotapi.InlineKeyboardButton, len(menu))
	for index, btn := range menu {
		fMenu[index] = tgbotapi.NewInlineKeyboardButtonData(btn.display, btn.command)
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
