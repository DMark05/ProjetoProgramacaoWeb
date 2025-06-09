package middleware

import (
	"net/http"
	"strings"
)

func RedirectDefaultWrongCallsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Compare(r.Header.Get("Content-Type"), "application/json") != 0 {
			http.Redirect(w, r, "http://localhost:3000/", http.StatusSeeOther)
		}
		next.ServeHTTP(w, r)
	})
}

func ValidateJWT(next http.Handler, validateFunc func(*string) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header["Authorization"][0]
		if err := validateFunc(&authToken); err != nil { //if fails
			http.Error(w, "No valid token", http.StatusUnauthorized)
		} else { //if succeeds
			next.ServeHTTP(w, r)
		}
	})
}
