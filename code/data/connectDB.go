// Package data represent connecting to Postgresql Database /*
package data

import (
	"database/sql"
	"fmt"
)

// ConnectDB represent connecting to Postgresql Database
func ConnectDB() *sql.DB {
	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgresdb", 5432, "bot_admin", "UbuLock!", "bot_admin")

	db, errdb := sql.Open("postgres", postgresInfo)
	if errdb != nil {
		panic(errdb)
	}

	return db
}
