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
	http.HandleFunc("/", middleware.OnlyPath("/", middleware.OnlyGet(handlers.IndexHandler)))
	http.HandleFunc("/register", middleware.OnlyPath("/", middleware.OnlyGet(handlers.RegisterHandler)))
	http.HandleFunc("/login", middleware.OnlyPath("/", middleware.OnlyGet(handlers.LoginHandler)))
	http.HandleFunc("/dashboard", middleware.OnlyPath("/", middleware.OnlyGet(handlers.DashBoard)))
	http.HandleFunc("/profile", middleware.OnlyPath("/", middleware.OnlyGet(handlers.MarketHandler)))
	http.HandleFunc("/ai-assistant", middleware.OnlyPath("/", middleware.OnlyGet(handlers.Ai_Assistant)))
	http.HandleFunc("/profile", middleware.OnlyPath("/", middleware.OnlyGet(handlers.ProfileHandler)))
}
