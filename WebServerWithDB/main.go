package main

import (
	"database-example/db"
	"database-example/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func startServer() {
	database := db.InitDB()
	if database == nil {
		log.Fatal("FAILED TO CONNECT TO DB")
	}

	router := mux.NewRouter().StrictSlash(true)

	encounterHandler := handler.NewEncounterHandler(database)
	encounterHandler.RegisterRoutes(router)

	encounterExecutionHandler := handler.NewEncounterExecutionHandler(database)
	encounterExecutionHandler.RegisterRoutes(router)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {

	startServer()
}
