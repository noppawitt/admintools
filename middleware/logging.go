package middleware

import (
	"log"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	Status int
	Length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.Status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.Status == 0 {
		w.Status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.Length += n
	return n, err
}

// LoggingMiddleware logs all incomming requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := statusWriter{ResponseWriter: w}
		next.ServeHTTP(&sw, r)
		duration := time.Now().Sub(start)
		log.Printf("%s %s %d %s - %d", r.Method, r.RequestURI, sw.Status, duration.String(), sw.Length)
	})
}
