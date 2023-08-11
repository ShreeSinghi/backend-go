package models

import (
	"fmt"
)

func Authenticate(cookieid string) (int, bool) {
	db, err := Connection()
	if err != nil {
		panic(err)
	}

	var userId int
	var admin bool

	err = db.QueryRow("SELECT userId FROM cookies WHERE cookies.sessionid = ?;", cookieid).Scan(&userId)

	if err != nil {
		panic(err)
	}

	err = db.QueryRow("SELECT admin FROM users WHERE id = ?;", userId).Scan(&admin)
	if err != nil {
		panic(err)
	}

	fmt.Println(userId, admin, cookieid)

	return userId, admin

}
