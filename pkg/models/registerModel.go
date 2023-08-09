package models

import (
	"database/sql"
	"errors"
	"log"
)

func RegisterUser(username, hash string) (int, error) {
	var userID int

	db, err := Connection()

	if err != nil {
		log.Fatal(err)
	}

	err = db.QueryRow(`SELECT id FROM users WHERE username = ?`, username).Scan(&userID)

	if err != nil {
		if err == sql.ErrNoRows {
			res, err := db.Exec(`INSERT INTO users (username, hash, admin) VALUES (?, ?, ?, 0)`, username, hash)
			if err != nil {
				return 0, err
			}

			userID, err := res.LastInsertId()
			if err != nil {
				return 0, err
			}

			return int(userID), nil
		}
		return 0, err
	}

	return 0, errors.New("user already exists")
}

func CreateCookie(userID int) error {
	_, err := db.Exec(`INSERT INTO cookies (userId) VALUES (?)`, userID)
	return err
}
