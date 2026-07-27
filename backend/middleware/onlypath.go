package middleware

import (
	"log"
	"net/http"
)

func HandlerFunc(path string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			log.Printf("%s %s", r.Method, r.URL)
			http.NotFound(w, r)
			return
		}
		next(w, r)
	}
}
