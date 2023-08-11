package controllers

import (
	"fmt"
	"log"
	"mvc/pkg/models"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func ProcessChecks(w http.ResponseWriter, r *http.Request) {

	//userId := r.Context().Value("userId").(int)
	admin := r.Context().Value("admin").(bool)

	db, err := models.Connection()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	r.ParseForm()

	checkRequests := r.PostForm
	delete(checkRequests, "admin")
	delete(checkRequests, "userId")

	if !admin {
		http.Error(w, "Not authenticated", http.StatusForbidden)
		return
	}

	for requestId := range checkRequests {
		var state string
		err := db.QueryRow("SELECT state FROM requests WHERE id = ?", requestId).Scan(&state)
		if err != nil {
			log.Fatal(err)
		}

		if state == "inrequested" {
			if checkRequests[requestId][0] == "approve" {
				_, err := db.Exec("UPDATE books SET quantity = quantity + 1 WHERE id = ?", requestId)
				if err != nil {
					log.Fatal(err)
				}

				_, err = db.Exec("DELETE FROM requests WHERE id = ?", requestId)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("returned")
			} else {
				_, err := db.Exec("UPDATE requests SET state = 'owned' WHERE id = ?", requestId)
				if err != nil {
					log.Fatal(err)
				}
			}
		} else {
			if checkRequests[requestId][0] == "approve" {
				_, err := db.Exec("UPDATE requests SET state='owned' WHERE id = ?", requestId)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(requestId, "approved")
			} else {
				bookIDStr := r.FormValue("bookId")
				bookID, err := strconv.Atoi(bookIDStr)
				if err != nil {
					log.Fatal(err)
				}

				_, err = db.Exec("UPDATE books SET quantity=quantity+1 WHERE id = ?", bookID)
				if err != nil {
					log.Fatal(err)
				}
				_, err = db.Exec("DELETE FROM requests WHERE id = ?", requestId)
				if err != nil {
					fmt.Println(requestId, "denied")
				}
			}
		}
	}
}
