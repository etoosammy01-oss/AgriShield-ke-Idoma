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
	statements := []string{
		`
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
                `,
		`
                CREATE TABLE IF NOT EXISTS crops (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        farmer_id INTEGER NOT NULL,
                        name TEXT NOT NULL,
                        quantity REAL NOT NULL DEFAULT 0,
                        unit TEXT NOT NULL,
                        location TEXT NOT NULL,
                        price_per_unit REAL NOT NULL DEFAULT 0,
                        listed_for_sale INTEGER NOT NULL DEFAULT 0,
                        image_url TEXT,
                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
                );
                `,
		`
                CREATE TABLE IF NOT EXISTS orders (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        buyer_id INTEGER NOT NULL,
                        crop_id INTEGER NOT NULL,
                        quantity REAL NOT NULL,
                        total_price REAL NOT NULL,
                        status TEXT NOT NULL DEFAULT 'completed',
                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
                );
                `,
		`
                CREATE TABLE IF NOT EXISTS diagnoses (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        farmer_id INTEGER NOT NULL,
                        image_name TEXT NOT NULL,
                        result TEXT NOT NULL,
                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
                );
                `,
		`
                CREATE TABLE IF NOT EXISTS negotiations (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        crop_id INTEGER NOT NULL,
                        buyer_id INTEGER NOT NULL,
                        farmer_id INTEGER NOT NULL,
                        quantity REAL NOT NULL,
                        status TEXT NOT NULL DEFAULT 'open',
                        round_count INTEGER NOT NULL DEFAULT 0,
                        max_rounds INTEGER NOT NULL DEFAULT 5,
                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                        expires_at DATETIME NOT NULL
                );
                `,
		`
                CREATE TABLE IF NOT EXISTS negotiation_messages (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        negotiation_id INTEGER NOT NULL,
                        sender_id INTEGER NOT NULL,
                        offer_price REAL NOT NULL,
                        message TEXT NOT NULL,
                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
                );
                `,
	}

	for _, statement := range statements {
		if _, err := db.Exec(statement); err != nil {
			return err
		}
	}

	log.Println("✅ Database tables migrated successfully")

	return nil
}
