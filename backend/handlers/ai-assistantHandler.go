package handlers

import (
	"backend/render"
	"log"
	"net/http"
)

func Ai_Assistant(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		//Text := r.FormValue("text")
	}
	log.Println("User Visited Ai Assistant page")
	if err := render.RenderTemplates(w, "ai-assistant.html", nil); err != nil {
		log.Fatalln("err Render Error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}