package handlers

import (
	"backend/render"
	"log"
	"net/http"
)

type FarmerDash struct {
	Goods     string
	Avaliable string
	Amount    string
}

func DashBoard(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("User Visited Dashboard")
		if err := render.RenderTemplates(w, "dashboard.html", nil); err != nil {
			log.Println("render error")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		log.Println("user's Choices")
	}
}
