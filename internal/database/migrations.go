package database

import "database/sql"

func RunMigrations(db *sql.DB) error {
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS links (
            id SERIAL PRIMARY KEY,
            original_url TEXT NOT NULL,
            short_code TEXT NOT NULL,
            clicks INT DEFAULT 0,
            created_at TIMESTAMP NOT NULL
        )
    `)
    return err
}