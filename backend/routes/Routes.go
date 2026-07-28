package routes

import (
	"backend/handlers"
	app "backend/internal"
	"backend/middleware"
	"net/http"
)

func RegisterRoutes(container *app.Container) {
	http.Handle(
		"/static/", http.StripPrefix("/static/",
			http.FileServer(http.Dir("../frontend")),
		),
	)
	http.HandleFunc("/", middleware.OnlyPath("/", middleware.OnlyGet(handlers.IndexHandler)))
	http.HandleFunc("/register", middleware.OnlyPath("/register", handlers.RegisterHandler))
	http.HandleFunc("/login", middleware.OnlyPath("/login", middleware.OnlyGet(handlers.LoginHandler)))
	http.HandleFunc("/dashboard", middleware.OnlyPath("/dashboard", middleware.OnlyGet(handlers.DashBoard)))
	http.HandleFunc("/profile", middleware.OnlyPath("/profile", middleware.OnlyGet(handlers.MarketHandler)))
	http.HandleFunc("/ai-assistant", middleware.OnlyPath("/ai-assistant", middleware.OnlyGet(handlers.Ai_Assistant)))
	//http.HandleFunc("/profile", middleware.OnlyPath("/", middleware.OnlyGet(handlers.ProfileHandler)))
}
