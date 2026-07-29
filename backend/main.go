package main

import (
	"log"
	"net/http"

	app "backend/internal"
	"backend/internal/database"
	"backend/internal/repository"
	"backend/routes"
)

func main() {

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := database.RunMigration(db); err != nil {
		log.Fatal(err)
	}

	// Create repository
	repo := repository.NewFarmerRepository(db)

	// Create container
	container := app.NewContainer(repo)

	routes.RegisterRoutes(container)

	log.Println("Server Starting on: http://localhost:8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
	    log.Println(err)
	}
}
