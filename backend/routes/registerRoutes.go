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
	http.HandleFunc("/register", middleware.Onlypath("/", middleware.OnlyGet(handlers.RegisterHandler)))
	http.HandleFunc("/login", middleware.Onlypath("/", middleware.OnlyGet(handlers.LoginHandler)))
	http.HandleFunc("/dashboard", middleware.Onlypath("/", middleware.OnlyGet(handlers.DashBoard)))
	http.HandleFunc("/profile", middleware.Onlypath("/", middleware.OnlyGet(handlers.MarketHandler)))
}
