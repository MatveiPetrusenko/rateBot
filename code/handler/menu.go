// Package handler serves for processing commands/keyboard menu/inline keyboard menu and for interactions with DB packages /*
package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/matthew/rateBot/configuration"
	"github.com/matthew/rateBot/data"
	"strconv"
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

// Greeting sending greeting message for new users
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

// ExchangeRate sending message with rate
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

// PlayGame represent navigation keyboard for playing the game
func PlayGame(bot *tgbotapi.BotAPI, chat *tgbotapi.Chat) {
	numericKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Easy"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Medium"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Show/Drop points"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Hide Menu"),
		),
	)
	msg := tgbotapi.NewMessage(chat.ID, "Choose difficult:")
	msg.ReplyMarkup = numericKeyboard
	bot.Send(msg)

	//checking record for current user exists or not
	result := data.CheckUser(chat.ID)

	switch result {
	case false:
		data.AddUser(chat.ID, chat.UserName) //creating record for collect statistic for new user
	default:
		return
	}
}

// EasyExercise sending easy level task and keyboard depends on progress user
func EasyExercise(bot *tgbotapi.BotAPI, chat *tgbotapi.Chat) {
	numericKeyboardEasy := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("I give-up! Give me one more, easy"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Back"),
		),
	)

	//getting string from DB or info about limit
	exerciseData := data.GetEasyResult(chat.ID)

	if exerciseData == "Exercises is over" {
		numericKeyboardBack := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Back"),
			),
		)

		msgTask := tgbotapi.NewMessage(chat.ID, exerciseData)
		msgTask.ReplyMarkup = numericKeyboardBack

		bot.Send(msgTask)
		return
	}

	msgEasy := tgbotapi.NewMessage(chat.ID, "Difficult: Easy")
	msgTask := tgbotapi.NewMessage(chat.ID, exerciseData)

	msgEasy.ReplyMarkup = numericKeyboardEasy
	bot.Send(msgEasy)
	bot.Send(msgTask)
}

// MediumExercise sending medium level task and keyboard depends on progress user
func MediumExercise(bot *tgbotapi.BotAPI, chat *tgbotapi.Chat) {
	numericKeyboardMedium := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("I give-up! Give me one more, medium"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Back"),
		),
	)

	//getting string from DB or info about limit
	exerciseData := data.GetMediumResult(chat.ID)

	if exerciseData == "Exercises is over" || exerciseData == "Do easy level to unlock medium" {
		numericKeyboardBack := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Back"),
			),
		)

		msgTask := tgbotapi.NewMessage(chat.ID, exerciseData)
		msgTask.ReplyMarkup = numericKeyboardBack

		bot.Send(msgTask)
		return
	}

	msgEasy := tgbotapi.NewMessage(chat.ID, "Difficult: Medium")
	msgTask := tgbotapi.NewMessage(chat.ID, exerciseData)

	msgEasy.ReplyMarkup = numericKeyboardMedium
	bot.Send(msgEasy)
	bot.Send(msgTask)
}

// NewEasyExercise represent sub keyboard for work on easy level task (New easy exercise/Back)
func NewEasyExercise() tgbotapi.ReplyKeyboardMarkup {
	numericKeyboardEasy := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("New easy exercise"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Back"),
		),
	)

	return numericKeyboardEasy
}

// NewMediumExercise represent sub keyboard for work on medium level task (New medium exercise/Back)
func NewMediumExercise() tgbotapi.ReplyKeyboardMarkup {
	numericKeyboardMedium := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("New medium exercise"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Back"),
		),
	)

	return numericKeyboardMedium
}

// Statistics serves for show results for user and represent inline keyboard (Yes/No)
func Statistics(bot *tgbotapi.BotAPI, chat *tgbotapi.Chat) {
	InlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Yes", "YES"),
			tgbotapi.NewInlineKeyboardButtonData("No", "NO"),
		),
	)

	total := data.UserScore(chat.ID)

	msgData := "You Total Score:" + strconv.Itoa(total) + "\n" + "Would you like to reset statistics and progress?"

	replyStatsStr := tgbotapi.NewMessage(chat.ID, msgData)

	replyStatsStr.ReplyMarkup = InlineKeyboard
	bot.Send(replyStatsStr)
}

// ResetStatistics dropping progress and score of user and sending message with notice
func ResetStatistics(bot *tgbotapi.BotAPI, chat *tgbotapi.Chat) {
	data.ResetUserScoreAndProgress(chat.ID)
	msgData := "Statistics and progress was reset"

	msg := tgbotapi.NewMessage(chat.ID, msgData)
	bot.Request(msg)
}

// GoalAndTotal return from DB total score and points for each solved task
func GoalAndTotal(userId int64) (int, int) {
	data.IncreaseUserProgress(userId)

	totalValue, exerciseValue := data.IncreaseUserScore(userId)

	return totalValue, exerciseValue
}

// CheckExceptionMenu return result of check on exceptions. Menu PlayGame (textMessage)
func CheckExceptionMenu(textMessage string) bool {
	exceptionString := map[string]bool{
		"Easy":                                true,
		"I give-up! Give me one more, easy":   true,
		"I give-up! Give me one more, medium": true,
		"Medium":                              true,
		"Show/Drop points":                    true,
		"Hide Menu":                           true,
		"Back":                                true,
		"New easy exercise":                   true,
		"New medium exercise":                 true,
	}

	_, ok := exceptionString[textMessage]
	if ok != true {
		return false
	} else {
		return true
	}
}

// CheckExceptionCommand return result of check on exceptions. Commands received from telegram (commandMessage)
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

// CheckExceptionDB return result of check on exceptions. Result from DB (textMessage)
func CheckExceptionDB(messageText string, userId int64) (bool, string) {
	resultCheckAnswer := data.CheckAnswer(messageText, userId)

	if resultCheckAnswer {
		return true, messageText
	} else {
		return false, ""
	}
}
