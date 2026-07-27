package handlers

import (
	"backend/render"
	"log"
	"net/http"
)
type FarmerDash struct {
	Goods string
	Avaliable string
	Amount string
}
func DashBoard(w http.ResponseWriter, r *http.Request) {
	log.Println("User Visited Dashboard page")
	if err := render.RenderTemplates(w, "dashboard.html", nil); err != nil {
		log.Println("error Render can't load Dashboard")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}