package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"log"
	"strconv"

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
	router.HandleFunc("/encounters/getById/{id}", h.GetByID).Methods("GET")
	router.HandleFunc("/encounters/getByCheckPoint/{id}", h.GetByCheckPointID).Methods("GET")

	router.HandleFunc("/encounters/delete/{id}", h.Delete).Methods("DELETE")

	router.HandleFunc("/encounters/update", h.Update).Methods("PUT")

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

func (handler *EncounterHandler) GetByID(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Encounter with ID: %d", id)
	encounter, err := handler.EncounterService.Find(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *EncounterHandler) GetByCheckPointID(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Encounter with ID: %d", id)
	encounter, err := handler.EncounterService.FindByCheckPointId(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *EncounterHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	// Dohvatite ID susreta iz URL parametara
	params := mux.Vars(req)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error parsing ID:", err)
	}

	// Pozovite servis za brisanje susreta
	err = handler.EncounterService.Delete(id)
	if err != nil {
		log.Println("Error deleting encounter:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Ako je brisanje uspješno, vratite status kod 204 No Content
	writer.WriteHeader(http.StatusNoContent)

}
func (handler *EncounterHandler) Update(writer http.ResponseWriter, req *http.Request) {
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
}
