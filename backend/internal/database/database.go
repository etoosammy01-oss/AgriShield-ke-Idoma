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
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err := db.Exec(farmersTable); err != nil {
		return err
	}

	// If this is an existing database created before "role" existed,
	// add the column now. SQLite has no "ADD COLUMN IF NOT EXISTS", so
	// we check first.
	if !hasColumn(db, "farmers", "role") {
		if _, err := db.Exec(`ALTER TABLE farmers ADD COLUMN role TEXT NOT NULL DEFAULT 'farmer'`); err != nil {
			return err
		}
	}

	cropsTable := `
	CREATE TABLE IF NOT EXISTS crops (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		farmer_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		quantity REAL NOT NULL,
		unit TEXT NOT NULL,
		location TEXT NOT NULL,
		price_per_unit REAL NOT NULL DEFAULT 0,
		listed_for_sale INTEGER NOT NULL DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (farmer_id) REFERENCES farmers(id)
	);
	`
	if _, err := db.Exec(cropsTable); err != nil {
		return err
	}

	ordersTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		buyer_id INTEGER NOT NULL,
		crop_id INTEGER NOT NULL,
		quantity REAL NOT NULL,
		total_price REAL NOT NULL,
		status TEXT NOT NULL DEFAULT 'completed',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (buyer_id) REFERENCES farmers(id),
		FOREIGN KEY (crop_id) REFERENCES crops(id)
	);
	`
	if _, err := db.Exec(ordersTable); err != nil {
		return err
	}

	diagnosesTable := `
	CREATE TABLE IF NOT EXISTS diagnoses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		farmer_id INTEGER NOT NULL,
		image_name TEXT NOT NULL,
		result TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (farmer_id) REFERENCES farmers(id)
	);
	`
	_, err := db.Exec(diagnosesTable)
	return err
}

func hasColumn(db *sql.DB, table, column string) bool {
	rows, err := db.Query("PRAGMA table_info(" + table + ")")
	if err != nil {
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var cid, notnull, pk int
		var name, ctype string
		var dflt sql.NullString
		if err := rows.Scan(&cid, &name, &ctype, &notnull, &dflt, &pk); err != nil {
			continue
		}
		if name == column {
			return true
		}
	}
	return false
}
