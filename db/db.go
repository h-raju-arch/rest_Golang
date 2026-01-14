package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB() {
	connStr := ""
	var err error
	Db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	pingerr := Db.Ping()
	if pingerr != nil {
		log.Fatal(pingerr)
	}
}

func CloseDb() {
	Db.Close()
}
