package handlers

import (
	"backend/render"
	"log"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("User Visited Profile")
	if err := render.RenderTemplates(w, "profile.html", nil); err != nil {
		log.Fatalln("err Render Problem", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}