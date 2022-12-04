package data

import "database/sql"

var connStr = "user=bot_admin password=UbuLock! dbname=bot_admin sslmode=disable"

func ConnectDB() *sql.DB {
	db, errdb := sql.Open("postgres", connStr)
	if errdb != nil {
		panic(errdb)
	}

	return db
}

//Need defer???
