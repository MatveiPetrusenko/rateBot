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

	result := data.CheckUser(chat.ID)

	switch result {
	case false:
		data.AddUser(chat.ID, chat.UserName)
	default:
		return
	}
}

// EasyExercise easy task for user
func EasyExercise(bot *tgbotapi.BotAPI, chat *tgbotapi.Chat) {
	numericKeyboardEasy := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("I give-up! Give me one more, easy"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Back"),
		),
	)

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

// MediumExercise medium task for user
func MediumExercise(bot *tgbotapi.BotAPI, chat *tgbotapi.Chat) {
	numericKeyboardMedium := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("I give-up! Give me one more, medium"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Back"),
		),
	)

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

// NewEasyExercise ...
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

// NewMediumExercise ...
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

// Statistics ...
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

// ResetStatistics ...
func ResetStatistics(bot *tgbotapi.BotAPI, chat *tgbotapi.Chat) {
	data.ResetUserScoreAndProgress(chat.ID)
	msgData := "Statistics and progress was reset"

	msg := tgbotapi.NewMessage(chat.ID, msgData)
	bot.Request(msg)
}

// GoalAndTotal ...
func GoalAndTotal(userId int64) (int, int) {
	data.IncreaseUserProgress(userId)

	totalValue, exerciseValue := data.IncreaseUserScore(userId)

	return totalValue, exerciseValue
}

// CheckExceptionMenu return result of check on exceptions (textMessage)
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

// CheckExceptionDB ...
func CheckExceptionDB(messageText string, userId int64) (bool, string) {
	resultCheckAnswer := data.CheckAnswer(messageText, userId)

	if resultCheckAnswer {
		return true, messageText
	} else {
		return false, ""
	}
}
