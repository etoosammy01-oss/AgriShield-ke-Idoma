package handlers

import (
	"backend/render"
	//"backend/internal2"
	"log"
	"net/http"
)
type MainMarket struct {
	ID string
	Name string
	Location string
	State string
}
func MarketHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("User Visited Market Page")
		if err := render.RenderTemplates(w, "market.html", nil); err != nil {
			log.Println("Render Error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
     User := MainMarket {
		ID: r.FormValue("id"),
		Name: r.FormValue("name"),
		Location: r.FormValue("location"),
		State: r.FormValue("state"),
	 }
	 
	 http.Redirect(w,r, User.ID, http.StatusSeeOther)
	}
}