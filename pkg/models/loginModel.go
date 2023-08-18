package models

import (
	"log"

	"github.com/google/uuid"
)

func Login(username string, password string) (string, error, bool) {
	db, err := Connection()

	if err != nil {
		log.Fatal(err)
	}

	var (
		id   int
		hash string
	)

	err = db.QueryRow(`SELECT id, hash FROM users WHERE username = (?)`, username).Scan(&id, &hash)

	if err != nil {
		log.Println("Error during login:", err)
	}

	if id == 0 || !MatchKaro(password, hash) {
		return "", nil, false
	}

	sessionString := uuid.New().String()
	_, err = db.Exec(`INSERT INTO cookies (userId, sessionID) VALUES (?, ?)`, id, sessionString)
	if err != nil {
		log.Fatal(err)
	}

	return sessionString, nil, true
}

func Logout(cookieid string) {
	db, err := Connection()
	_, err = db.Exec("DELETE FROM cookies WHERE sessionId = ?", cookieid)
	if err != nil {
		log.Fatal(err)
	}
}
