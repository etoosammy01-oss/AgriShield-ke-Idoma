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
	log.Println("User Just visit the registration")
	if err := render.RenderTemplates(w, "register.html", nil); err != nil {
		log.Println("render can't load html file", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}