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

type EncounterHandler struct {
	EncounterService *service.EncounterService
}

func NewEncounterHandler(db *gorm.DB) *EncounterHandler {
	encounterService := service.NewEncounterService(db)
	return &EncounterHandler{EncounterService: encounterService}
}

func (h *EncounterHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/encounters", h.Create).Methods("POST")
}

func (handler *EncounterHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var encounter model.Encounter

	// Ispisi telo zahtjeva prije nego što se pokuša dekodirati JSON
	body, errr := ioutil.ReadAll(req.Body)
	fmt.Println("errr", errr)
	fmt.Println("Primljeno telo zahtjeva:", string(body))

	errs := json.Unmarshal(body, &encounter)
	if errs != nil {
		fmt.Println("Greška pri parsiranju JSON-a:", errs)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Ispis mape
	fmt.Println("Mapa nakon parsiranja JSON-a:", encounter)

	errp := handler.EncounterService.Create(&encounter)
	if errp != nil {
		println("Error while creating a new encounter")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
