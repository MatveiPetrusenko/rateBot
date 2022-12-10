package data

import (
	"database/sql"
	"fmt"
)

// MaxIdValue ...
func MaxIdValue() (int, int) {
	db := ConnectDB()

	defer db.Close()

	var maxIdEasy, maxIdMedium int

	requestEasyTable := "SELECT exercise_id FROM easy_economic_exercises WHERE exercise_id = (SELECT MAX (exercise_id) FROM easy_economic_exercises);"
	sqlStatementDataE := db.QueryRow(requestEasyTable)
	errE := sqlStatementDataE.Scan(&maxIdEasy)
	if errE != nil {
		if errE == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(errE)
		}
	}

	requestMediumTable := "SELECT exercise_id FROM medium_economic_exercises WHERE exercise_id = (SELECT MAX (exercise_id) FROM medium_economic_exercises);"
	sqlStatementDataM := db.QueryRow(requestMediumTable)
	errM := sqlStatementDataM.Scan(&maxIdMedium)
	if errM != nil {
		if errM == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(errM)
		}
	}

	return maxIdEasy, maxIdMedium
}

// UserProgress ...
func UserProgress(userId int64) int {
	db := ConnectDB()

	defer db.Close()

	var userProgressValue int

	request := "SELECT progress FROM user_results WHERE user_id = $1"

	sqlStatementData := db.QueryRow(request, userId)
	errH := sqlStatementData.Scan(&userProgressValue)
	if errH != nil {
		if errH == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(errH)
		}
	}

	return userProgressValue
}
