package middleware

import (
	"net/http"
)

func StripSlashes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var path string = r.URL.Path
		if len(path) > 1 && path[len(path)-1] == '/' {
			r.URL.Path = path[:len(path)-1]
		}

		next.ServeHTTP(w, r)
	})
}