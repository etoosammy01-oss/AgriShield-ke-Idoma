package handlers

import (
	"backend/render"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("User Visited Home Page")
	if err := render.RenderTemplates(w, "home.html", nil); err != nil {
		log.Println("Render can not Load template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
