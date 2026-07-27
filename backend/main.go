package main

import (
	"backend/routes"
	"log"
	"net/http"
)

func main() {
	routes.RegisterRoutes()
	log.Println("Server Activated on: http//localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
	}
}
