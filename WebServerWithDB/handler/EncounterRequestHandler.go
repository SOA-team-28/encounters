package handler

import (
	"database-example/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type EncounterRequestHandler struct {
	EncounterRequestService *service.EncounterRequestService
}

func NewEncounterRequestHandler(db *gorm.DB) *EncounterRequestHandler {
	encounterRequestService := service.NewEncounterRequestService(db)
	return &EncounterRequestHandler{EncounterRequestService: encounterRequestService}
}
func (h *EncounterRequestHandler) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/requests", h.GetAll).Methods("GET")
	router.HandleFunc("/requests/accept/{id}", h.Accept).Methods("PUT")
	router.HandleFunc("/requests/reject/{id}", h.Reject).Methods("PUT")

}
func (handler *EncounterRequestHandler) GetAll(writer http.ResponseWriter, req *http.Request) {

	requests, err := handler.EncounterRequestService.FindAll()
	if err != nil {
		fmt.Println("Greška pri dohvatu requests:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(requests)
	if err != nil {
		fmt.Println("Greška pri enkodiranju JSON odgovora:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("Odgovor: ")
	fmt.Println(requests)
}
func (handler *EncounterRequestHandler) Accept(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EncounterRequestService.Accept(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
func (handler *EncounterRequestHandler) Reject(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EncounterRequestService.Reject(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
