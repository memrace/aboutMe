package commands

import (
	"errors"
	"unicode/utf8"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"aboutMe/dialogs"
)

const creatorChatId int64 = 282568630

const readMe = "readme"

const start = "start"

const faq = "faq"

const contactMe = "contactme"

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

	if dialog := dialogs.GetDialog(handler.update.Message.From.ID); dialog != nil && !dialog.Replied {
		if reply := message.Text; utf8.RuneCountInString(reply) > 0 {
			dialog.Reply = reply
			dialog.Replied = true
			dialogs.SaveDialog(dialog)
			sendMessage(handler.bot, makeMessage(message.Chat.ID, "Я отправил сообщение хозяину\nКак только получу ответ вернусь к вам!"))
			sendMessage(handler.bot, makeMessage(creatorChatId, "Сообщение от "+"\nпользователя: "+dialog.FirstName+" "+dialog.LastName+"\n"+dialog.Reply))
		} else {
			sendMessage(handler.bot, makeMessage(message.Chat.ID, "Bad reply :)"))
		}
	} else {
		if message.IsCommand() {
			workWithCommand(handler.bot, handler.update)
		} else {
			workWithSimpleMessage(message.Chat.ID, handler.bot)
		}
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

	case contactMe:
		go func() {
			user := message.From
			dialogs.SaveDialog(&dialogs.Dialog{
				UserID:    user.ID,
				UserName:  user.UserName,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				ChatID:    message.Chat.ID,
			})
			sendMessage(bot, makeMessage(message.Chat.ID, contactMeResponse))
		}()
		return
	default:
		go workWithSimpleMessage(message.Chat.ID, bot)
		return
	}
}

func workWithSimpleMessage(chatID int64, bot *tgbotapi.BotAPI) {
	sendMessage(bot, makeMessage(chatID, "¯\\_(ツ)_/¯"))
}

const contactMeResponse = "Укажите ваше имя, компанию и желаемую дату и время \nФормат: Имя Компания Дата Время"
