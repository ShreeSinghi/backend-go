package controllers

import (
	"mvc/pkg/models"

	"context"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie := r.Header.Get("Cookie")
		if len(cookie) < 10 {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			cookieid := r.Header.Get("Cookie")[10:]
			userId, admin := models.Authenticate(cookieid)

			ctx := context.WithValue(r.Context(), "userId", userId)
			ctx = context.WithValue(ctx, "admin", admin)

			// Call the next handler with the updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
