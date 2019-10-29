package middleware

import "net/http"

type Middleware struct {}

// COMMON
func (m *Middleware) COMMON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}


// CORS
func (m *Middleware) CORS(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		next.ServeHTTP(w, r)
	})
}

func InitMiddleware() *Middleware {
	return &Middleware{}
}