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
	EncounterService               *service.EncounterService
	HiddenLocationEncounterService *service.HiddenLocationEncounterService
	SocialEncounterService         *service.SocialEncounterService
	EncounterRequestService        *service.EncounterRequestService
}

func NewEncounterHandler(db *gorm.DB) *EncounterHandler {
	encounterService := service.NewEncounterService(db)
	hiddenLocationEncounterService := service.NewHiddenLocationEncounterService(db)
	socialEncounterService := service.NewSocialEncounterService(db)
	encounterRequestService := service.NewEncounterRequestService(db)
	return &EncounterHandler{
		EncounterService:               encounterService,
		HiddenLocationEncounterService: hiddenLocationEncounterService,
		SocialEncounterService:         socialEncounterService,
		EncounterRequestService:        encounterRequestService,
	}
}

func (h *EncounterHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/encounters", h.Create).Methods("POST")
	router.HandleFunc("/encounters/createByTourist/{checkpointId}/{touristId}", h.CreateEncounterByTourist).Methods("POST")
	router.HandleFunc("/encounters/getById/{id}", h.GetByID).Methods("GET")
	router.HandleFunc("/encounters/getByCheckPoint/{id}", h.GetByCheckPointID).Methods("GET")

	router.HandleFunc("/encounters/delete/{id}", h.Delete).Methods("DELETE")

	router.HandleFunc("/encounters/update", h.Update).Methods("PUT")
	router.HandleFunc("/encounters", h.GetAll).Methods("GET")

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

func (handler *EncounterHandler) GetAll(writer http.ResponseWriter, req *http.Request) {

	encounters, err := handler.EncounterService.FindAll()
	if err != nil {
		fmt.Println("Greška pri dohvatu encounters:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(encounters)
	if err != nil {
		fmt.Println("Greška pri enkodiranju JSON odgovora:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("Odgovor: ")
	fmt.Println(encounters)
}

func (handler *EncounterHandler) CreateEncounterByTourist(writer http.ResponseWriter, req *http.Request) {
	var encounter model.Encounter
	params := mux.Vars(req)
	idCheckPointString := params["checkpointId"]
	idTouristString := params["touristId"]
	checkpointId, err := strconv.Atoi(idCheckPointString)
	if err != nil {
		fmt.Println("Greška pri citanju parametara i checkpointId-a:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	touristId, err := strconv.Atoi(idTouristString)

	fmt.Println("Procitani id check pointa", checkpointId)
	if err != nil {
		fmt.Println("Greška pri citanju parametara i checkpointId-a:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Ispisi telo zahtjeva prije nego što se pokuša dekodirati JSON
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Greška pri čitanju zahtjeva:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("Primljeno telo zahtjeva:", string(body))
	err = json.Unmarshal(body, &encounter)
	if err != nil {
		fmt.Println("Greška pri parsiranju JSON-a:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	var createdEncounter interface{}
	encounter.CheckPointId = checkpointId
	switch encounter.Type {
	case "Location":
		createdEncounter = handler.HiddenLocationEncounterService.CreateHiddenLocationEncounter(&encounter)
		createdEncounter = handler.EncounterService.Create(&encounter)
	case "Social":
		createdEncounter = handler.SocialEncounterService.CreateSocialEncounterByTourist(&encounter)
		createdEncounter = handler.EncounterService.Create(&encounter)
	case "Mics":
		createdEncounter = handler.EncounterService.Create(&encounter)
	default:
		fmt.Println("Greška pri kreiranju susreta, tip nepoznat!", err)
	}

	var request model.EncounterRequest

	allEncounters, err := handler.EncounterService.EncounterRepo.FindAll()
	if err != nil {
		fmt.Println("Greška pri parsiranju JSON-a:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	request.EncounterId = len(allEncounters)
	request.Status = "OnHold"
	request.TouristId = touristId

	err = handler.EncounterRequestService.Create(&request)

	if err != nil {
		println("Error while creating a new request")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	// Vratite kreirani susret kao JSON odgovor
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(createdEncounter)
	if err != nil {
		fmt.Println("Greška pri enkodiranju JSON odgovora:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Kreiran susret:", createdEncounter)
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

	// Ispis mape
	fmt.Println("Mapa nakon parsiranja JSON-a:", encounter)

	errp := handler.EncounterService.Update(&encounter)
	if errp != nil {
		println("Error while creating a new encounter")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
