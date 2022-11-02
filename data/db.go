/*
Package data provides functionality for connect to DataBase
*/
package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"math/rand"
)

var ConnStr = "user=postgres password=postgres dbname=exercise_database sslmode=disable"

// GetResult ...
func GetResult() (string, string) {
	db, errdb := sql.Open("postgres", ConnStr)
	if errdb != nil {
		panic(errdb)
	}

	defer db.Close()

	var header, body string

	min := 1
	max := 10
	randomValue := min + rand.Intn(max-min+1)

	sqlStatementHeader := db.QueryRow("SELECT header FROM easy_exercise WHERE task_id = $1", randomValue)
	errH := sqlStatementHeader.Scan(&header)
	if errH != nil {
		if errH == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(errH)
		}
	}

	sqlStatementBody := db.QueryRow("SELECT task_desc FROM easy_exercise WHERE task_id = $1", randomValue)
	errB := sqlStatementBody.Scan(&body)
	if errB != nil {
		if errB == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(errB)
		}
	}

	return header, body
}
