package models

import (
	"log"
)

func Authenticate(cookieid string) (int, bool) {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}

	var userId int
	var admin bool

	err = db.QueryRow("SELECT users.id, users.admin FROM users, cookies WHERE cookies.sessionid = (?);", cookieid).Scan(&userId, &admin)
	if err != nil {
		log.Println(cookieid)
		log.Fatal(err)
	}
	return userId, admin

}
