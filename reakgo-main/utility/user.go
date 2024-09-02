// utility/user.go
package utility

import (
    "log"
    "github.com/jmoiron/sqlx"
)

// InsertUser inserts a new user into the database
func InsertUser(db *sqlx.DB, name, email, address, password string) error {
    query := `
        INSERT INTO users (name, email, address, password)
        VALUES (?, ?, ?, ?)
    `
    _, err := db.Exec(query, name, email, address, password)
    if err != nil {
        log.Println("Error inserting user:", err)
        return err
    }
    return nil
}
