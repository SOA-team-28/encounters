package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=super dbname=gorm port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Encounter{})
	return database
}

func startServer(handler *handler.EncounterHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/encounters", handler.Create).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	eRepo := &repo.EncounterRepository{DatabaseConnection: database}
	eService := &service.EncounterService{EncounterRepo: eRepo}
	eHandler := &handler.EncounterHandler{EncounterService: eService}

	startServer(eHandler)
}
