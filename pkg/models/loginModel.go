package models

import (
	"crypto/rand"
	"errors"
	"log"
)

func Login(username string, password string) ([]byte, error) {
	db, err := Connection()

	if err != nil {
		log.Fatal(err)
	}

	var (
		id   int
		salt string
		hash string
	)

	err = db.QueryRow("SELECT id, salt, hash FROM users WHERE username = $1", username).Scan(&id, &salt, &hash)

	if err != nil {
		log.Println("Error during login:", err)
	}

	if id == 0 || !MatchKaro(password, hash) {
		return []byte{}, errors.New("invalid username or password")
	}

	newSessionID := make([]byte, 16)
	_, err = rand.Read(newSessionID)
	if err != nil {
		log.Fatal(err)
	}

	return newSessionID, nil
}
