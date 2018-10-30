package middleware

import "net/http"

// Middleware is a base middleware type
type Middleware func(http.Handler) http.Handler
