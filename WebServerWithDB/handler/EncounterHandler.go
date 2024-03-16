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

	// Parsirajte JSON telo u mapu
	var data map[string]interface{}
	erre := json.Unmarshal(body, &data)
	if erre != nil {
		fmt.Println("Greška pri parsiranju JSON-a:", erre)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Ispis mape
	fmt.Println("Mapa nakon parsiranja JSON-a:", data)

	errs := json.Unmarshal(body, &encounter)
	if errs != nil {
		fmt.Println("Greška pri parsiranju JSON-a:", errs)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Ispis mape
	fmt.Println("Mapa nakon parsiranja JSON-a:", encounter)
	/*
		err := json.NewDecoder(req.Body).Decode(&encounter)
		if err != nil {

			println("Error while parsing json")
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
	*/
	errp := handler.EncounterService.Create(&encounter)
	if errp != nil {
		println("Error while creating a new encounter")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
