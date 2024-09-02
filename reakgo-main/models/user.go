package models

import (
    "log"
    "reakgo/utility" // Adjust import path
)

// Define the User model
type User struct {
	Id   int
    Name     string
    Email    string
    Address  string
    Password string
}

// InsertUser inserts a user into the database
func InsertUser(user User) error {
    query := `
        INSERT INTO users (name, email, address, password)
        VALUES (?, ?, ?, ?)`
    _, err := utility.Db.Exec(query, user.Name, user.Email, user.Address, user.Password)
    if err != nil {
        log.Println("Database insert error:", err)
        return err
    }
    return nil
}
