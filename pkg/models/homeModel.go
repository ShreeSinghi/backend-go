package models

type Book struct {
	ID       int   
	Title    string
	Quantity int    
}

type Request struct {
	ID     int 
	bookId int    
	UserId int    
	State  string
	Title  string 

}

type User struct {
	ID        int  
	Username  string 
	Requested bool   
	Admin     bool
}

type Cookie struct {
	ID        int  
	UserId    int    
	SessionID string 
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
			return nil, err
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
			return nil, err
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
					ownedBooks = append(ownedBooks, book)
					break
				}
			}
		}
	}

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
			if request.bookId == book.ID {
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
	rows, err = db.Query("SELECT id, username, requested FROM users WHERE requested = true")
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
		"books":    booksResult,
		"booksout": outList,
		"booksin":  inList,
		"users":    usersResult,
	}

	return data, nil
}
