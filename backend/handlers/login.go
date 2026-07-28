package handlers

import (
	"backend/render"
	"log"
	"net/http"
)

type LoginData struct {
	Phone    string
	Password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("User Visited Login Page")
		if err := render.RenderTemplates(w, "login.html", nil); err != nil {
			log.Println("Render Error")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		user := LoginData{
			Phone:    r.FormValue("phone"),
			Password: r.FormValue("password"),
		}
		if user.Phone == "" || user.Password == "" {
			log.Println("Both Phone Number And Password Are Required")
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	default:
		http.Error(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
	}

}
