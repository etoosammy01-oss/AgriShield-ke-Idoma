package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./agro-shield.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("✅ Connected to SQLite database")

	return db, nil
}

func RunMigration(db *sql.DB) error {
	farmersTable := `
	CREATE TABLE IF NOT EXISTS farmers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		full_name TEXT NOT NULL,
		phone TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		location TEXT NOT NULL,
		role TEXT NOT NULL DEFAULT 'farmer',
		photo_url TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(farmersTable)
	if err != nil {
		return err
	}

	log.Println("✅ Farmers table migrated successfully")

	return nil
}
