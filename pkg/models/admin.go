package models

import (
	"database/sql"
	"strconv"
)

func AddBook(title string, quantity int) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	var (
		id          int
		oldquantity int
	)

	err = db.QueryRow("SELECT id, quantity FROM books WHERE title = (?)", title).Scan(&id, &oldquantity)

	if err == sql.ErrNoRows {
		_, err := db.Exec("INSERT INTO books (title, quantity) VALUES (?, ?)", title, quantity)
		if err != nil {
			return err
		}

	} else if err != nil {
		return err

	} else {
		_, err := db.Exec("UPDATE books SET quantity = quantity + ? WHERE title = ?", quantity, title)
		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessChecks(checkRequests map[string][]string) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	for requestId := range checkRequests {
		var state string
		var userId int
		err := db.QueryRow("SELECT state, userId FROM requests WHERE id = ?", requestId).Scan(&state, &userId)
		if err != nil {
			return err
		}

		if state == "inrequested" {
			if checkRequests[requestId][0] == "approve" {
				_, err := db.Exec("UPDATE books SET quantity = quantity + 1 WHERE id = ?", requestId)
				if err != nil {
					return err
				}

				_, err = db.Exec("DELETE FROM requests WHERE id = ?", requestId)
				if err != nil {
					return err
				}

			} else {
				_, err := db.Exec("UPDATE requests SET state = 'owned' WHERE id = ?", requestId)
				if err != nil {
					return err
				}
			}
		} else {
			if checkRequests[requestId][0] == "approve" {
				_, err := db.Exec("UPDATE requests SET state='owned' WHERE id = ?", requestId)
				if err != nil {
					return err
				}
			} else {
				bookIDStr := requestId // Assuming the request ID is the book ID
				bookID, err := strconv.Atoi(bookIDStr)
				if err != nil {
					return err
				}

				_, err = db.Exec("UPDATE books SET quantity=quantity+1 WHERE id = ?", bookID)
				if err != nil {
					return err
				}

				_, err = db.Exec("DELETE FROM requests WHERE id = ?", requestId)
				if err != nil {
					return err
				}
			}
			_, err = db.Exec("UPDATE users SET requested = false WHERE id = ?", userId)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func ProcessAdminRequests(requestedUsers map[string][]string) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	for userID, action := range requestedUsers {
		if action[0] == "approve" {
			var username string
			err := db.QueryRow("SELECT username FROM users WHERE id = ? AND requested = true", userID).Scan(&username)

			if err == sql.ErrNoRows {
				return err
			}

			if err != nil {
				return err
			}

			_, err = db.Exec("UPDATE users SET admin = true, requested = false WHERE id = ?", userID)
			if err != nil {
				return err
			}
		} else {
			var userId int
			err := db.QueryRow("SELECT * FROM users WHERE id = ? AND requested = true", userID).Scan(&userId)

			if err == sql.ErrNoRows {
				return err
			}

			if err != nil {
				return err
			}

			_, err = db.Exec("UPDATE users SET requested = false WHERE id = ?", userId)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
