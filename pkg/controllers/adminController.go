package controllers

import (
	"fmt"
	"log"
	"mvc/pkg/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	router.POST("/process-checks", processChecks)

	router.Run(":8080")
}

func processChecks(c *gin.Context) {

	db, err = models.Connection()
	admin := c.PostForm("admin")
	checkRequests := c.Request.PostForm
	delete(checkRequests, "admin")
	delete(checkRequests, "userId")

	if admin == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "Not authenticated",
		})
		return
	}

	for requestId := range checkRequests {
		var state string
		err := db.QueryRow("SELECT state FROM requests WHERE id = ?", requestId).Scan(&state)
		if err != nil {
			log.Fatal(err)
		}

		if state == "inrequested" {
			if checkRequests[requestId] == "approve" {
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
			if checkRequests[requestId] == "approve" {
				_, err := db.Exec("UPDATE requests SET state='owned' WHERE id = ${db.escape(requestId)}")
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(requestId, "apprived")
			} else {
				fmt.Println(results)
				_, err = db.Exec("UPDATE books SET quantity=quantity+1 WHERE id = ${results[0].bookId}")
				if err != nil {
					log.Fatal(err)
				}
				_, err = db.Exec("DELETE FROM requests WHERE id = ${db.escape(requestId)}")
				if err != nil {
					fmt.Println(requestId, "denied")

				}
			}
		}
	}
}
