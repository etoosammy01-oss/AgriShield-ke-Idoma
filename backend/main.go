package main

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/repository"
	"backend/routes"
	"log"
	"net/http"
)

func main() {

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.RunMigration(db); err != nil {
		log.Fatal(err)
	}

	log.Println("Server Starting on: http://localhost:8080 ...")

	routes.RegisterRoutes()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
	}

	farmerRepo := repository.NewFarmerRepository(db)

	farmer := &models.Farmer{
		FullName:     "Test Farmer",
		Phone:        "08011112222",
		PasswordHash: "hashed-password",
		Location:     "Otukpo",
	}

	err = farmerRepo.Create(farmer)
	if err != nil {
		log.Println("Insert failed:", err)
	} else {
		log.Println("Farmer inserted successfully")
	}
}
