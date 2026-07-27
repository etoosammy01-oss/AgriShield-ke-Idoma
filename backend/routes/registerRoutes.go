package routes

import (
	"backend/handlers"
	"backend/middleware"
	"net/http"
)

func RegisterRoutes() {
	http.Handle(
		"/static/", http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	http.HandleFunc("/", middleware.Onlypath("/", middleware.OnlyGet(handlers.IndexHandler)))
}
