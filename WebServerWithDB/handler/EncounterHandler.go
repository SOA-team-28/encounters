package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type EncounterHandler struct {
	EncounterService *service.EncounterService
}

func (handler *EncounterHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var encounter model.Encounter

	// Ispisi telo zahtjeva prije nego što se pokuša dekodirati JSON
	body, errr := ioutil.ReadAll(req.Body)
	fmt.Println("errr", errr)
	fmt.Println("Primljeno telo zahtjeva:", string(body))

	err := json.NewDecoder(req.Body).Decode(&encounter)
	if err != nil {

		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EncounterService.Create(&encounter)
	if err != nil {
		println("Error while creating a new encounter")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
