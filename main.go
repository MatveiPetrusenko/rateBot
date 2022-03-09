/*
package main provides functionality for ...
*/
package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/matthew/rateBot/cbr_api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5147742910:AAGzy-TUkxpuLkEKej0kh2Z3gQJ1jICQgns")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	//ctx, _ := context.WithTimeout(context.Background(), 3*time.Minute)

	//timer := time.NewTimer(time.Second)

	/*	sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT)*/
	for {
		select {
		/*		case s := <-sigs:
				fmt.Println(s)*/

		/*		case <-timer.C:
				fmt.Println("Tick")
				timer.Reset(time.Second)*/

		/*		case <-ctx.Done():
				fmt.Println("Bye!")
				break loop*/

		case update := <-updates:
			UserName := update.Message.From.UserName
			UserID := update.Message.From.ID
			ChatID := update.Message.Chat.ID
			Text := update.Message.Text

			//print user message
			log.Printf("[%s] %d %d %s", UserName, UserID, ChatID, Text)

			//special commands
			switch update.Message.Command() {
			case "start":
				greeting(bot, update.Message)
			case "exchangerate":
				exchangeRate(bot, update.Message)
			case "playgame":
				playGame(bot, update.Message.Chat)
			case "getintouch":
				getintouch(bot, update.Message)
			}

			greetingDescr := "Please select want you want\n" +
				"You can check rate or play the game"
			msg := tgbotapi.NewMessage(ChatID, greetingDescr)

			switch Text {
			case "Easy":
				easyExercise(bot, update.Message.Chat)
			case "I give-up! Give me one more, easy":
				easyExercise(bot, update.Message.Chat)
			case "I give-up! Give me one more, medium":
				mediumExercise(bot, update.Message.Chat)
			case "Medium":
				mediumExercise(bot, update.Message.Chat)
			case "Hide Menu":
				msg.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{RemoveKeyboard: true}
				bot.Send(msg)
			case "Back":
				msg.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{RemoveKeyboard: true}
				bot.Send(msg)
				playGame(bot, update.Message.Chat)
			}

			//replying on undesignated message
			/*			reply := "Unidentified message.\n" + "Try again."
						msg := tgbotapi.NewMessage(ChatID, reply)
						bot.Send(msg)*/
		}

	}
}

// Grerting do things...
func greeting(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	UserName := message.From.UserName

	greetingStr := "Hello" + "\xF0\x9F\x91\x8B" + "\n" + UserName
	greetingDescr := "Please select want you want\n" +
		"You can check rate or play the game"

	var gretingArray = [2]string{greetingStr, greetingDescr}

	for i := 0; i < 2; i++ {
		greetingMsg := tgbotapi.NewMessage(message.Chat.ID, gretingArray[i])
		bot.Send(greetingMsg)
	}
}

// Getintouch do things...
func getintouch(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	getintouchStr := "You can text to him: @ma1thew"

	getintouchMsg := tgbotapi.NewMessage(message.Chat.ID, getintouchStr)
	bot.Send(getintouchMsg)
}

//exchangerate do things
func exchangeRate(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	replyRateStr := tgbotapi.NewMessage(message.Chat.ID, cbr_api.ParsingJSON())
	bot.Send(replyRateStr)
}

// ExchangeRate do things...
func playGame(bot *tgbotapi.BotAPI, chatID *tgbotapi.Chat) {
	numericKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Easy"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Medium"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Hide Menu"),
		),
	)
	msg := tgbotapi.NewMessage(chatID.ID, "Choose difficult:")
	msg.ReplyMarkup = numericKeyboard
	bot.Send(msg)

}

// EasyExercise do things...
func easyExercise(bot *tgbotapi.BotAPI, chatID *tgbotapi.Chat) {
	numericKeyboardEasy := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("I give-up! Give me one more, easy"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Back"),
		),
	)

	msgEasy := tgbotapi.NewMessage(chatID.ID, "Difficult: Easy")
	msgEasy.ReplyMarkup = numericKeyboardEasy
	bot.Send(msgEasy)

	// out of order
	/*	first, second := task_data.GetResult()

		exMsgF := tgbotapi.NewMessage(chatID.ID, first)
		bot.Send(exMsgF)

		exMsgS := tgbotapi.NewMessage(chatID.ID, second)
		bot.Send(exMsgS)*/
}

// MediumExercise do things...
func mediumExercise(bot *tgbotapi.BotAPI, chatID *tgbotapi.Chat) {
	numericKeyboardMedium := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("I give-up! Give me one more, medium"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Back"),
		),
	)

	msgMedium := tgbotapi.NewMessage(chatID.ID, "Difficult: Medium")
	msgMedium.ReplyMarkup = numericKeyboardMedium
	bot.Send(msgMedium)

	exStr := "Some text, medium exercise"
	exMsg := tgbotapi.NewMessage(chatID.ID, exStr)
	bot.Send(exMsg)
}

/*//todo
replace simvol $
correct parsind json date
separate main on several files
add comand in decs.*/

/*func OnMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	// Пользователь, который написал боту
	userName := message.From.UserName

	// ID чата/диалога.
	// Может быть идентификатором как чата с пользователем
	// (тогда он равен UserID) так и публичного чата/канала
	chatID := message.Chat.ID

	log.Printf("[%s] %d", userName, chatID)

	if update.Message.Ne.UserName != ""
}*/
