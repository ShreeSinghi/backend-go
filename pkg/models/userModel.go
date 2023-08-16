package models

import (
	"fmt"
)

func RequestCheckout(bookId int, userId int) (interface{}, error) {
	db, err := Connection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var bookQuantity int
	err = db.QueryRow("SELECT quantity FROM books WHERE id = ?", bookId).Scan(&bookQuantity)
	if err != nil {
		return nil, err
	}
	if bookQuantity == 0 {
		return nil, fmt.Errorf("Book is out of stock")
	}

	var existingRequest int
	err = db.QueryRow("SELECT COUNT(*) FROM requests WHERE bookId = ? AND userId = ?", bookId, userId).Scan(&existingRequest)
	if err != nil {
		return nil, err
	}
	if existingRequest > 0 {
		return GetDataUser(userId, "You have already requested this book")
	}

	_, err = db.Exec("INSERT INTO requests (bookId, userId, state) VALUES (?, ?, 'outrequested')", bookId, userId)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("UPDATE books SET quantity = quantity - 1 WHERE id = ?", bookId)
	if err != nil {
		return nil, err
	}

	return GetDataUser(userId, "Checkout request submitted")
}

func RequestCheckin(bookId int, userId int) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	var existingRequestID int
	err = db.QueryRow("SELECT id FROM requests WHERE bookId = ? AND userId = ? AND state = 'owned'", bookId, userId).Scan(&existingRequestID)
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE requests SET state = 'inrequested' WHERE id = ?", existingRequestID)
	if err != nil {
		return err
	}

	return nil
}

func RequestAdmin(userId int) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	fmt.Println(userId)

	_, err = db.Exec("UPDATE users SET requested = true WHERE id = ?", userId)
	if err != nil {
		return err
	}

	return nil
}
