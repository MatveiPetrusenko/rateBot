// Package data provides functionality for work with DB /*
package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// GetEasyResult return data string from table (easy_economic_exercises) or limit string if exercise is not exist
func GetEasyResult(userId int64) string {
	db := ConnectDB()

	defer db.Close()

	var dataExercise string
	var exerciseId int

	//getting max id exercise
	maxIdEasy, _ := MaxIdValue()

	//check user table
	userProgressValue := UserProgress(userId)

	if userProgressValue >= maxIdEasy {
		return "Exercises is over"
	}

	switch userProgressValue {
	case 0:
		exerciseId = 1
	default:
		exerciseId = userProgressValue + 1
	}

	requestData := "SELECT data_exercise FROM easy_economic_exercises WHERE exercise_id = $1"

	sqlStatementData := db.QueryRow(requestData, exerciseId)
	errH := sqlStatementData.Scan(&dataExercise)
	if errH != nil {
		if errH == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(errH)
		}
	}

	return dataExercise
}

// GetMediumResult return data string from table (medium_economic_exercises) or limit string if exercise is not exist
func GetMediumResult(userId int64) string {
	db := ConnectDB()

	defer db.Close()

	var dataExercise string
	var exerciseId int

	//getting max id exercise
	maxIdEasy, maxIdMedium := MaxIdValue()

	//check user table
	userProgressValue := UserProgress(userId)

	if userProgressValue < maxIdEasy {
		return "Do easy level to unlock medium"
	}

	//limit
	if userProgressValue >= maxIdMedium {
		return "Exercises is over"
	}

	exerciseId = userProgressValue + 1

	requestData := "SELECT data_exercise FROM medium_economic_exercises WHERE exercise_id = $1"

	sqlStatementData := db.QueryRow(requestData, exerciseId)
	errH := sqlStatementData.Scan(&dataExercise)
	if errH != nil {
		if errH == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(errH)
		}
	}

	return dataExercise
}

// AddUser create new row for new user in table (user_results)
func AddUser(id int64, userName string) {
	db := ConnectDB()

	defer db.Close()

	request := "INSERT INTO user_results VALUES ($1,$2,$3, $4);"
	if _, err := db.Exec(request, id, "@"+userName, 0, 0); err != nil {
		fmt.Println(err)
	}
}

// CheckUser check table (user_results), record exist ot not
func CheckUser(id int64) bool {
	db := ConnectDB()

	defer db.Close()

	var responceId int64

	request := "SELECT user_id FROM user_results WHERE user_id = $1"
	sqlStatementData := db.QueryRow(request, id)
	errH := sqlStatementData.Scan(&responceId)
	if errH != nil {
		if errH == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(errH)
		}
	}

	if id == responceId {
		return true
	} else {
		return false
	}
}

// IncreaseUserProgress increase on +1 progress of user in table (user_results)
func IncreaseUserProgress(userId int64) {
	db := ConnectDB()

	defer db.Close()

	request := "UPDATE user_results SET progress = progress + 1 WHERE user_id = $1;"
	if _, err := db.Exec(request, userId); err != nil {
		fmt.Println(err)
	}
}

// CheckAnswer checking for the correctness of the answer in one of the tables (easy_economic_exercises/medium_economic_exercises)
func CheckAnswer(textMessage string, userId int64) bool {
	db := ConnectDB()

	defer db.Close()

	var dataAnswer string

	//getting user progress
	userProgressValue := UserProgress(userId)

	//getting max id value
	maxIdEasy, _ := MaxIdValue()

	switch userProgressValue {
	case 0:
		userProgressValue = 1
	default:
		userProgressValue += 1
	}

	//var tableValue = "easy_economic_exercises"
	requestData := "SELECT answer FROM easy_economic_exercises WHERE exercise_id = $1;"
	//
	if userProgressValue >= maxIdEasy {
		//tableValue = "medium_economic_exercises"
		requestData = "SELECT answer FROM medium_economic_exercises WHERE exercise_id = $1;"
	}

	//requestData := "SELECT answer FROM easy_economic_exercises WHERE exercise_id = $1;"

	sqlStatementData := db.QueryRow(requestData, userProgressValue)
	errH := sqlStatementData.Scan(&dataAnswer)
	if errH != nil {
		if errH == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(errH)
		}
	}

	if textMessage != dataAnswer {
		return false
	} else {
		return true
	}
}

// IncreaseUserScore transmits score from tables (easy_economic_exercises/medium_economic_exercises) for each solved exercise in different table (user_results)
func IncreaseUserScore(userId int64) (int, int) {
	db := ConnectDB()

	defer db.Close()

	//getting user progress
	progressValue := UserProgress(userId)

	//getting max id value
	maxIdEasy, maxIdMedium := MaxIdValue()

	request := "UPDATE user_results SET score_point = score_point + (SELECT score_point FROM easy_economic_exercises WHERE exercise_id = $1) WHERE user_id = $2;"

	if progressValue <= maxIdEasy {
		request = "UPDATE user_results SET score_point = score_point + (SELECT score_point FROM easy_economic_exercises WHERE exercise_id = $1) WHERE user_id = $2;"
	} else if maxIdEasy < progressValue && progressValue <= maxIdMedium {
		request = "UPDATE user_results SET score_point = score_point + (SELECT score_point FROM medium_economic_exercises WHERE exercise_id = $1) WHERE user_id = $2;"
	}

	if _, err := db.Exec(request, progressValue, userId); err != nil {
		fmt.Println(err)
	}

	var scoreData int

	request = "SELECT score_point FROM user_results WHERE user_id = $1"

	if progressValue <= maxIdEasy {
		request = "SELECT score_point FROM easy_economic_exercises WHERE exercise_id = $1"
	} else if maxIdEasy < progressValue && progressValue <= maxIdMedium {
		request = "SELECT score_point FROM medium_economic_exercises WHERE exercise_id = $1"
	}

	sqlStatementData := db.QueryRow(request, progressValue)
	errH := sqlStatementData.Scan(&scoreData)
	if errH != nil {
		if errH == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(errH)
		}
	}

	request = "SELECT score_point FROM user_results WHERE user_id = $1"

	var totalScore int

	sqlStatementDataU := db.QueryRow(request, userId)
	errHU := sqlStatementDataU.Scan(&totalScore)
	if errHU != nil {
		if errHU == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(errHU)
		}
	}

	return totalScore, scoreData
}

// UserScore getting total score for user from table (user_results)
func UserScore(userId int64) int {
	db := ConnectDB()

	defer db.Close()

	request := "SELECT score_point FROM user_results WHERE user_id = $1"

	var totalScore int

	sqlStatementData := db.QueryRow(request, userId)
	errH := sqlStatementData.Scan(&totalScore)
	if errH != nil {
		if errH == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(errH)
		}
	}

	return totalScore
}

// ResetUserScoreAndProgress reset user score and progress in table (user_results)
func ResetUserScoreAndProgress(userId int64) {
	db := ConnectDB()

	defer db.Close()

	request := "UPDATE user_results SET score_point = 0, progress = 0 WHERE user_id = $1;"
	if _, err := db.Exec(request, userId); err != nil {
		fmt.Println(err)
	}
}
