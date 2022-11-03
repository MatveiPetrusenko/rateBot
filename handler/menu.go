package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/matthew/rateBot/configuration"
)

var InlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇺🇸", "US"),
		tgbotapi.NewInlineKeyboardButtonData("🇪🇺", "EC"),
		tgbotapi.NewInlineKeyboardButtonData("🇬🇧", "UK"),
		tgbotapi.NewInlineKeyboardButtonData("🇨🇭", "CH"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇷🇺", "RU"),
		tgbotapi.NewInlineKeyboardButtonData("🇺🇦", "UA"),
		tgbotapi.NewInlineKeyboardButtonData("🇧🇾", "BY"),
		tgbotapi.NewInlineKeyboardButtonData("🇰🇿", "KZ"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Hide", "Hide"),
	),
)

// Greeting sending greeting message for new user
func Greeting(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	userName := message.From.UserName

	greetingStr := "Hello" + "\xF0\x9F\x91\x8B" + "\n" + userName
	greetingDescr := "Please select want you want\n" +
		"You can check rate or play the game"

	var gretingArray = [2]string{greetingStr, greetingDescr}

	for i := 0; i < 2; i++ {
		greetingMsg := tgbotapi.NewMessage(message.Chat.ID, gretingArray[i])
		bot.Send(greetingMsg)
	}
}

// ExchangeRate sending rate
func ExchangeRate(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	replyRateStr := tgbotapi.NewMessage(message.Chat.ID, configuration.ParsingJSON())

	replyRateStr.ReplyMarkup = InlineKeyboard
	bot.Send(replyRateStr)
}

// GetInTouch sending tg contact
func GetInTouch(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	getintouchStr := "You can text to him: @ma1thew"

	getintouchMsg := tgbotapi.NewMessage(message.Chat.ID, getintouchStr)
	bot.Send(getintouchMsg)
}

// PlayGame run keyboard
func PlayGame(bot *tgbotapi.BotAPI, chatID *tgbotapi.Chat) {
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

// EasyExercise easy task for user
func EasyExercise(bot *tgbotapi.BotAPI, chatID *tgbotapi.Chat) {
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
	/*	first, second := data.GetResult()

		exMsgF := tgbotapi.NewMessage(chatID.ID, first)
		bot.Send(exMsgF)

		exMsgS := tgbotapi.NewMessage(chatID.ID, second)
		bot.Send(exMsgS)*/
}

// MediumExercise medium task for user
func MediumExercise(bot *tgbotapi.BotAPI, chatID *tgbotapi.Chat) {
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

// CheckExceptionMenu return result of check on exceptions (textMessage)
func CheckExceptionMenu(textMessage string) bool {
	exceptionString := map[string]bool{
		"Easy":                                true,
		"I give-up! Give me one more, easy":   true,
		"I give-up! Give me one more, medium": true,
		"Medium":                              true,
		"Hide Menu":                           true,
		"Back":                                true,
	}

	_, ok := exceptionString[textMessage]
	if ok != true {
		return false
	} else {
		return true
	}
}

// CheckExceptionCommand return result of check on exceptions (commandMessage)
func CheckExceptionCommand(botCommand string) bool {
	exceptionString := map[string]bool{
		"start":        true,
		"exchangerate": true,
		"playgame":     true,
		"getintouch":   true,
	}

	_, ok := exceptionString[botCommand]
	if ok != true {
		return false
	} else {
		return true
	}
}
