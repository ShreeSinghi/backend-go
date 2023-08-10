package models

type Book struct {
	ID       int    //
	Title    string //`db:"title"`
	Quantity int    //`db:"quantity"`
}

type Request struct {
	ID     int    // `db:"id"`
	BookID int    //`db:"book_id"`
	UserID int    //`db:"user_id"`
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
	UserID    int    //`db:"userId"`
	SessionID string //`db:"sessionId"`
}

func GetDataUser(userID int, checkoutStatus string) (interface{}, error) {
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
		booksResult = append(booksResult, book)
	}

	var requestsResult []Request
	rows, err = db.Query("SELECT * FROM requests WHERE userId=(?)", userID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var request Request
		err := rows.Scan(&request.ID, &request.BookID, &request.UserID, &request.State)
		if err != nil {
			panic(err)
		}
		requestsResult = append(requestsResult, request)
	}

	var userresult User
	err = db.QueryRow("SELECT * FROM users WHERE id=(?)", userID).Scan(&userresult.ID, &userresult.Username, &userresult.Requested)
	if err != nil {
		return nil, err
	}

	ownedBooks := make([]*Book, 0)
	for _, request := range requestsResult {
		if request.State == "owned" || request.State == "inrequested" {
			for _, book := range booksResult {
				if book.ID == request.BookID {
					ownedBooks = append(ownedBooks, book)
					break
				}
			}
		}
	}

	data := map[string]interface{}{
		"books":            booksResult,
		"isadminrequested": userresult.Requested,
		"ownedbooks":       ownedBooks,
		"checkoutStatus":   checkoutStatus,
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
	rows, err := db.Query("SELECT * FROM requests WHERE state IN ('outrequested', 'inrequested')")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var request Request
		err := rows.Scan(&request.ID, &request.UserID, &request.BookID, &request.State)
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
		booksResult = append(booksResult, book)
	}

	outList := make([]Request, 0)
	inList := make([]Request, 0)

	for _, request := range requestsResult {
		for _, book := range booksResult {
			if request.BookID == book.ID {
				request.Title = book.Title
				break
			}
		}
		if request.State == "outrequested" {
			outList = append(outList, request)
		} else if request.State == "inrequested" {
			inList = append(inList, request)
		}
	}

	usersResult := make([]User, 0)
	rows, err = db.Query("SELECT * FROM users WHERE requested = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Requested)
		if err != nil {
			return nil, err
		}
		usersResult = append(usersResult, user)
	}

	data := map[string]interface{}{
		"booksout": outList,
		"booksin":  inList,
		"users":    usersResult,
	}

	return data, nil
}
