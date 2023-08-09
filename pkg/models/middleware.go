package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("sessionID")
		if err != nil {
			http.Error(w, "Not authenticated", http.StatusForbidden)
			return

		}

		sessionId := cookie.Value
		if strings.Contains(r.Header.Get("Cookie"), "sessionID") {

			var userId int
			var admin bool
			var username string

			err := db.QueryRow(`
                SELECT cookies.userId, cookies.sessionId, users.admin, users.username
                FROM cookies
                JOIN users ON cookies.userId = users.id
                WHERE cookies.sessionId = ?`,
				sessionId).Scan(&userId, &sessionId, &admin, &username)

			if err != nil || sessionId != cookie.Value {

				http.Error(w, "not authenticated", http.StatusForbidden)
				return
			}

			r.Header.Set("userId", fmt.Sprint(userId))
			r.Header.Set("admin", fmt.Sprint(admin))
			r.Header.Set("username", username)
		} else {

			http.Error(w, "not authenticated", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}
