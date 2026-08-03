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

	farmerRepo := repository.NewFarmerRepository(db)
	cropRepo := repository.NewCropRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	diagnosisRepo := repository.NewDiagnosisRepository(db)
	negotiationRepo := repository.NewNegotiationRepository(db)
	negotiationMsgRepo := repository.NewNegotiationMessageRepository(db)

	container := app.NewContainer(farmerRepo, cropRepo, orderRepo, diagnosisRepo, negotiationRepo, negotiationMsgRepo)

	routes.RegisterRoutes(container)

	log.Println("Server Starting on: http://localhost:9000 ...")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Println(err)
	}
}
