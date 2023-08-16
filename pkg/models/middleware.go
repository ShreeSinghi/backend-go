package models

func Authenticate(cookieid string) (int, bool, bool) {
	db, err := Connection()
	if err != nil {
		panic(err)
	}

	var userId int
	var admin bool
	var authorised bool = true

	err = db.QueryRow("SELECT userId FROM cookies WHERE cookies.sessionid = ?;", cookieid).Scan(&userId)

	if err != nil {
		authorised = false
	}

	err = db.QueryRow("SELECT admin FROM users WHERE id = ?;", userId).Scan(&admin)
	if err != nil {
		authorised = false
	}

	return userId, admin, authorised

}
