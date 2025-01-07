package db

import (
	"database/sql"
	"log"
	"os"
)
const dbPath = "sqlite-database.db"

func InitDB() {

	os.Remove(dbPath)
	file, err := os.Create(dbPath)
	file.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error()5553wz)
	}
}