package middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware logs all incomming requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method + " " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
