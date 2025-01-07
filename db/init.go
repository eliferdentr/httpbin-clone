package db

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite" // Modern SQLite driver
)

const dbPath = "sqlite-database.db"

var DB *sql.DB

func InitDB() {
	var err error
	// Veritabanı bağlantısını aç
	DB, err = sql.Open("sqlite", dbPath) // Modern SQLite driver'ı kullan
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if err =DB.Ping(); err !=nil {
		log.Fatal(err.Error())
	}

	// Bağlantı havuz ayarları
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Tabloları oluştur
	createTables()
}

func createTables() {
	createTableString := `CREATE TABLE IF NOT EXISTS request_logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		method TEXT NOT NULL,
		endpoint TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(createTableString)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	log.Println("Tables created successfully")
}
