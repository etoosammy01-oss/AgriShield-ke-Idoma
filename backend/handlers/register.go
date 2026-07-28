package handlers

import (
	"backend/render"
	"log"
	"net/http"
)
type UserReg struct{
	Name string
	Phone string
	Email string
	Password string
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet :
		log.Println("User Visited Register page")
		if err := render.RenderTemplates(w, "register.html", nil); err != nil {
			log.Println("render error", err)
			http.Error(w, "Internal Server error", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		log.Println("form submitted successfully")
		http.Redirect(w,r, "/login", http.StatusSeeOther)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}