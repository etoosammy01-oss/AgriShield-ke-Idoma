package handlers

import (
	"backend/render"
	"log"
	"net/http"
)
	type LoginData struct {
		Phone string
		Password string
	}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("user Visited Login Page")
	if err := render.RenderTemplates(w, "login.html", nil); err != nil {
		log.Println("err render can not load login html")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}