package models

import (
	"fmt"
	"log"
)

type Book struct {
	ID       int    //
	Title    string //`db:"title"`
	Quantity int    //`db:"quantity"`
}

type Request struct {
	ID     int    // `db:"id"`
	bookId int    //`db:"book_id"`
	UserId int    //`db:"user_id"`
	State  string //`db:"state"`
	Title  string // not unpacked by db directly

}

type User struct {
	ID        int    //`db:"id"`
	Username  string // `db:"username"`
	Requested bool   // `db:"requested"`
	Admin     bool
}

type Cookie struct {
	ID        int    // `db:"id"`
	UserId    int    //`db:"UserId"`
	SessionID string //`db:"sessionId"`
}

func GetDataUser(userId int, checkoutStatus string) (interface{}, error) {
	db, err := Connection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	booksResult := make([]*Book, 0)
	rows, err := db.Query("SELECT * FROM books")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		book := new(Book)
		err := rows.Scan(&book.ID, &book.Title, &book.Quantity)
		if err != nil {
			panic(err)
		}

		if book.Quantity > 0 {
			booksResult = append(booksResult, book)
		}

	}

	var requestsResult []Request
	rows, err = db.Query("SELECT * FROM requests WHERE userId=(?)", userId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var request Request
		err := rows.Scan(&request.ID, &request.bookId, &request.UserId, &request.State)
		if err != nil {
			panic(err)
		}
		requestsResult = append(requestsResult, request)
	}

	var userresult User
	err = db.QueryRow("SELECT id, username, requested FROM users WHERE id=(?)", userId).Scan(&userresult.ID, &userresult.Username, &userresult.Requested)
	if err != nil {
		return nil, err
	}

	ownedBooks := make([]*Book, 0)
	for _, request := range requestsResult {
		if request.State == "owned" || request.State == "inrequested" {
			for _, book := range booksResult {
				if book.ID == request.bookId {
					log.Println(book)
					ownedBooks = append(ownedBooks, book)
					break
				}
			}
		}
	}

	fmt.Println(booksResult[0])

	data := map[string]interface{}{
		"books":          booksResult,
		"userData":       userresult,
		"ownedBooks":     ownedBooks,
		"checkoutStatus": checkoutStatus,
	}

	return data, nil
}

func GetDataAdmin() (interface{}, error) {

	db, err := Connection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var requestsResult []Request
	rows, err := db.Query("SELECT id, userId, bookId, state FROM requests WHERE state IN ('outrequested', 'inrequested')")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var request Request
		err := rows.Scan(&request.ID, &request.UserId, &request.bookId, &request.State)
		if err != nil {
			return nil, err
		}
		requestsResult = append(requestsResult, request)
	}

	var booksResult []Book
	rows, err = db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Quantity)
		if err != nil {
			return nil, err
		}
		if book.Quantity > 0 {
			booksResult = append(booksResult, book)
		}
	}

	outList := make([]Request, 0)
	inList := make([]Request, 0)

	for _, request := range requestsResult {
		for _, book := range booksResult {
			fmt.Println(request.bookId, book.ID)
			if request.bookId == book.ID {
				request.Title = book.Title
				break
			}
		}
		fmt.Println(request.Title, request.UserId, request.State)
		if request.State == "outrequested" {
			outList = append(outList, request)
		} else if request.State == "inrequested" {
			inList = append(inList, request)
		}
	}

	usersResult := make([]User, 0)
	rows, err = db.Query("SELECT id, username, requested FROM users WHERE requested = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		fmt.Println("meow")
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Requested)
		if err != nil {
			return nil, err
		}
		usersResult = append(usersResult, user)
	}

	log.Println(booksResult)

	data := map[string]interface{}{
		"books":    booksResult,
		"booksout": outList,
		"booksin":  inList,
		"users":    usersResult,
	}

	return data, nil
}
