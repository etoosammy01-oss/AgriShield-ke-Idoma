package handlers

import (
	"backend/internal/app"
	"backend/render"
	"log"
	"net/http"
)


type UserReg struct {
	First_Name       string
	Last_Name        string
	Phone            string
	Email            string
	Password         string
	Confirm_Password string
}

func (h *Register) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("User Visited Register page")
		if err := render.RenderTemplates(w, "register.html", nil); err != nil {
			log.Println("render error", err)
			http.Error(w, "Internal Server error", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		user := UserReg{
			First_Name:       r.FormValue("first-name"),
			Last_Name:        r.FormValue("last-name"),
			Phone:            r.FormValue("phone"),
			Email:            r.FormValue("email"),
			Password:         r.FormValue("password"),
			Confirm_Password: r.FormValue("confirm-password"),
		}
		if user.First_Name == "" || user.Last_Name == "" || user.Phone == "" || user.Password == "" || user.Confirm_Password == "" {
			log.Println("user details must not be empty")
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		} else if user.Password != user.Confirm_Password {
			log.Println("Password Mismatch")
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		location := r.FormValue("location")

		err := app.AuthService.Register(
			user.First_Name,
			user.Last_Name,
			user.Phone,
			user.Password,
			location,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("Farmer registered successfully")

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
