package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"log"
	"strconv"

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
	router.HandleFunc("/executions/updateStatusByCheckPoint/{id}", h.UpdateStatusByCheckPointId).Methods("PUT")

	// DOHVAT SVIH ZAVRŠENIH ENCOUNTERA ZA TURISTU
	router.HandleFunc("/executions/get-all-completed/{touristID}", h.GetAllCompletedByTourist).Methods("GET")
}

func (handler *EncounterExecutionHandler) GetAllCompletedByTourist(writer http.ResponseWriter, req *http.Request) {
	// Izvući touristID iz zahtjeva
	vars := mux.Vars(req)
	fmt.Println("USaoooooo...")
	touristID, err := strconv.ParseInt(vars["touristID"], 10, 64)
	if err != nil {
		fmt.Println("Greška pri parsiranju touristID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Pozvati servis za dohvat završenih susreta za turistu
	encounterExecutions, err := handler.EncounterExecutionService.GetAllCompletedByTourist(touristID)
	if err != nil {
		fmt.Println("Greška pri dohvatu završenih susreta:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Vratiti završene susrete kao JSON odgovor
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(encounterExecutions)
	if err != nil {
		fmt.Println("Greška pri enkodiranju JSON odgovora:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("Odgovor: ")
	fmt.Println(encounterExecutions)
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

func (handler *EncounterExecutionHandler) UpdateStatusByCheckPointId(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Encounter with ID: %d", id)
	err = handler.EncounterExecutionService.UpdateStatus(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)

}
