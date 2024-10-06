package middlewares

import (
	"net/http"
)

func CorsMiddlewares(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		
		allowedOrigins := []string{
			"http://localhost:3000",
		}

		for _, origns := range allowedOrigins {
			if origin == origns {
				w.Header().Set("Access-Control-Allow-Origin", origns)
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Authorization")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}