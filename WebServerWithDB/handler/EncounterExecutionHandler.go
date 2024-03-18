package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type EncounterExecutionHandler struct {
	EncounterExecutionService *service.EncounterExecutionService
}

func NewEncounterExecutionHandler(db *gorm.DB) *EncounterExecutionHandler {
	encounterExecutionService := service.NewEncounterExecutionService(db)
	return &EncounterExecutionHandler{EncounterExecutionService: encounterExecutionService}
}

func (h *EncounterExecutionHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/executions/activate", h.Create).Methods("POST")

}

func (handler *EncounterExecutionHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var encounterExecution model.EncounterExecution

	// Ispisi telo zahtjeva prije nego što se pokuša dekodirati JSON
	body, errr := ioutil.ReadAll(req.Body)
	fmt.Println("errr", errr)
	fmt.Println("Primljeno telo zahtjeva:", string(body))

	errs := json.Unmarshal(body, &encounterExecution)
	if errs != nil {
		fmt.Println("Greška pri parsiranju JSON-a:", errs)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Ispis mape
	fmt.Println("Mapa nakon parsiranja JSON-a:", encounterExecution)

	errp := handler.EncounterExecutionService.Create(&encounterExecution)
	if errp != nil {
		println("Error while creating a new encounterExecution")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
