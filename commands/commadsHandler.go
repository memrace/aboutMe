package commands

import (
	"context"
	"errors"
	"unicode/utf8"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	service "aboutMe/api"
)

const creatorChatId int64 = 282568630

const readMe = "readme"

const start = "start"

const faq = "faq"

const contactMe = "contactme"

func MakeCommandHandler(
	service *service.ApiService,
	bot *tgbotapi.BotAPI,
	update *tgbotapi.Update,
) CommandHandler {
	return &commandHandler{
		Service: service,
		bot:     bot,
		update:  update,
	}
}

type CommandHandler interface {
	Process(ctx context.Context)
	ProcessCallback()
}

type commandHandler struct {
	Service *service.ApiService
	bot     *tgbotapi.BotAPI
	update  *tgbotapi.Update
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

func (handler *commandHandler) Process(ctx context.Context) {
	message := handler.update.Message

	chatId := message.Chat.ID
	userId := handler.update.Message.From.ID
	client := handler.Service.Client

	dialog, _ := client.Get(ctx, &service.GetDialog{
		UserId: userId,
	})
	if dialog != nil && !dialog.Replied {
		if reply := message.Text; utf8.RuneCountInString(reply) > 0 {

			_, err := client.SetReply(ctx, &service.UserReply{
				UserId: userId,
				Text:   reply,
			})
			if err != nil {
				sendErrorMessage(handler.bot, chatId, err)
			} else {
				sendMessage(handler.bot, makeMessage(chatId, "Я отправил сообщение создателю\nКак только получу ответ вернусь к вам!"))
				sendMessage(handler.bot, makeMessage(creatorChatId, "Сообщение от "+"\nпользователя: "+dialog.FirstName+" "+dialog.LastName+" "+dialog.UserName+"\n"+reply))

				_, err = client.Delete(ctx, &service.DialogId{
					Id: userId,
				})
				if err != nil {
					sendErrorMessage(handler.bot, chatId, err)
				}
			}

		} else {
			sendMessage(handler.bot, makeMessage(chatId, "Bad reply :)"))
		}
	} else {
		if message.IsCommand() {
			workWithCommand(ctx, handler.bot, handler.update, client)
		} else {
			workWithSimpleMessage(chatId, handler.bot)
		}
	}
}

func sendErrorMessage(bot *tgbotapi.BotAPI, chatId int64, err error) {
	sendMessage(bot, makeMessage(chatId, "Ошибка :( "+err.Error()))
	print(err)
}

func workWithCommand(ctx context.Context, bot *tgbotapi.BotAPI, update *tgbotapi.Update, client service.DialogServiceClient) {
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
			_, err := client.Create(ctx, &service.CreateDialog{
				Id:        user.ID,
				UserName:  user.UserName,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				ChatId:    message.Chat.ID,
			})
			if err != nil {
				sendErrorMessage(bot, message.Chat.ID, err)
			} else {
				sendMessage(bot, makeMessage(message.Chat.ID, contactMeResponse))
			}

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
