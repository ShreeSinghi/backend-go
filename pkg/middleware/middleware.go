package middleware

import (
	"mvc/pkg/models"

	"context"
	"net/http"
	"strings"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie := r.Header.Get("Cookie")
		if len(cookie) < 10 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			cookieid := cookie[strings.Index(cookie, "sessionID=")+10:]
			userId, admin, authorised := AuthenticateDB(cookieid)

			if !authorised {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			ctx := context.WithValue(r.Context(), "userId", userId)
			ctx = context.WithValue(ctx, "admin", admin)
			ctx = context.WithValue(ctx, "authorised", authorised)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}

func AuthenticateDB(cookieid string) (int, bool, bool) {
	db, err := models.Connection()
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
