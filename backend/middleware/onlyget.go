package middleware

import (
	"log"
	"net/http"
)

func HandleFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Println("Method Not Allowed")
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		next(w, r)

	}
}
