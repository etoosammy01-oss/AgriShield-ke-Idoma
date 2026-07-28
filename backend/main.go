package main

import (
	app "backend/internal"
	"backend/internal/database"
	"backend/internal/repository"
	"backend/internal/services"
	"backend/routes"
	"log"
	"net/http"
)

func main() {

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewFarmerRepository(db)

	container := app.NewContainer(repo)

	farmerRepo := repository.NewFarmerRepository(db)

	authService := services.NewAuthService(farmerRepo)
	_ = authService

	if err := database.RunMigration(db); err != nil {
		log.Fatal(err)
	}

	log.Println("Server Starting on: http://localhost:8080 ...")

	routes.RegisterRoutes(container)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
	}
}
