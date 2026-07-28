package handlers

import (
	"backend/render"
	"log"
	"net/http"
)
type MainMarket struct {
	
}
func MarketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("User visited Market Page")
	if err := render.RenderTemplates(w, "market.html", nil); err != nil {
		log.Fatalln("Render Not Loaded", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}