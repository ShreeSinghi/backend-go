package controllers

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
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			cookieid := cookie[strings.Index(cookie, "sessionID=")+10:]
			userId, admin := models.Authenticate(cookieid)

			ctx := context.WithValue(r.Context(), "userId", userId)
			ctx = context.WithValue(ctx, "admin", admin)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
