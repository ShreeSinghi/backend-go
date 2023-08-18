package models

import (
	"database/sql"
	"errors"
	"log"
)

func RegisterUser(username, hash string) (int, error) {
	var userId int

	db, err := Connection()

	if err != nil {
		log.Fatal(err)
	}

	err = db.QueryRow(`SELECT id FROM users WHERE username = ?`, username).Scan(&userId)

	if err != sql.ErrNoRows {
		if err != nil {
			return 0, err
		}
		return 0, errors.New("user already exists")
	} else {
		admin := 0

		// check if this is the first user then make them admin
		err = db.QueryRow(`SELECT id FROM users`).Scan(&userId)
		if err == sql.ErrNoRows {
			admin = 1
		}

		res, err := db.Exec(`INSERT INTO users (username, hash, admin) VALUES (?, ?, ?)`, username, hash, admin)
		if err != nil {
			return 0, err
		}

		userId, err := res.LastInsertId()
		if err != nil {
			return 0, err
		}
		return int(userId), nil
	}
}