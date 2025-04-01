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
