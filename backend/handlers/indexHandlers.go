package handlers

import (
	"backend/render"
	"log"
	"net/http"
)

type AgroDash struct {
	Name string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("User Visited Home Page")
	PageName := AgroDash{
		Name: "Agro Shield",
	}
	if err := render.RenderTemplates(w, "index.html", PageName); err != nil {
		log.Println("Render can not Load template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
