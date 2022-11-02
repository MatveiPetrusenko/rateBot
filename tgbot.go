/*
package main provides functionality for ...
*/
package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/matthew/rateBot/configuration"
	"github.com/matthew/rateBot/handler"
	"log"
)

// telegramBot connect by token
func telegramBot() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	bot, err := tgbotapi.NewBotAPI("5147742910:AAGzy-TUkxpuLkEKej0kh2Z3gQJ1jICQgns")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	return bot, updates
}

func main() {
	bot, updates := telegramBot()

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %d %d %s", update.Message.From.UserName, update.Message.From.ID, update.Message.Chat.ID, update.Message.Text) //print user message

			// check exception
			if handler.CheckExceptionMenu(update.Message.Text) == false && handler.CheckExceptionCommand(update.Message.Command()) == false {
				reply := "Non understandable message.\n" + "Try again." //replying on non designated message
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				bot.Send(msg)
			} else {

				//special commands
				switch update.Message.Command() {
				case "start":
					handler.Greeting(bot, update.Message)
				case "exchangerate":
					handler.ExchangeRate(bot, update.Message)
				case "playgame":
					handler.PlayGame(bot, update.Message.Chat)
				case "getintouch":
					handler.GetInTouch(bot, update.Message)
				}

				//handler navigation
				switch update.Message.Text {
				case "Easy":
					handler.EasyExercise(bot, update.Message.Chat)
				case "I give-up! Give me one more, easy":
					handler.EasyExercise(bot, update.Message.Chat)
				case "I give-up! Give me one more, medium":
					handler.MediumExercise(bot, update.Message.Chat)
				case "Medium":
					handler.MediumExercise(bot, update.Message.Chat)
				case "Hide Menu":
					greetingDescr := "Please select want you want\n" +
						"You can check rate or play the game"
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, greetingDescr)

					msg.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{RemoveKeyboard: true}
					bot.Send(msg)
				case "Back":
					greetingDescr := "Please select want you want\n" +
						"You can check rate or play the game"
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, greetingDescr)

					msg.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{RemoveKeyboard: true}
					bot.Send(msg)
					handler.PlayGame(bot, update.Message.Chat)
				}

			}
		} else if update.CallbackQuery != nil {
			var replyStr string

			switch update.CallbackQuery.Data {
			case "US":
				replyStr = "🇺🇸"
			case "EC":
				replyStr = "🇪🇺"
			case "UK":
				replyStr = "🇬🇧"
			case "CH":
				replyStr = "🇨🇭"
			case "RU":
				replyStr = "🇷🇺"
			case "UA":
				replyStr = "🇺🇦"
			case "BY":
				replyStr = "🇧🇾"
			case "KZ":
				replyStr = "🇰🇿"
			default:
				replyStr = "Hide"
			}

			// Respond to the callback query, telling Telegram to show the user a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, replyStr)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			emptyKeyBoard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("", "_")))

			editLastMessage := tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, emptyKeyBoard)

			if replyStr == "Hide" {
				bot.Request(editLastMessage)
			} else {

				bot.Request(editLastMessage)

				msgStr := configuration.SubMessage(replyStr)

				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, msgStr)
				msg.ReplyMarkup = handler.InlineKeyboard

				bot.Request(msg)
			}
		}
	}
}

/*//todo
correct parsind json date
add date in message
add comand in decs
*/
