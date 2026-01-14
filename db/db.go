package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB() {
	connStr := "postgresql://neondb_owner:npg_aeikcVqw2s9v@ep-dry-frost-ah8ww6ko-pooler.c-3.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"
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
