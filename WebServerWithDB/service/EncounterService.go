package service

import (
	"database-example/model"
	"database-example/repo"

	"fmt"

	"gorm.io/gorm"
)

type EncounterService struct {
	EncounterRepo *repo.EncounterRepository
}

func NewEncounterService(db *gorm.DB) *EncounterService {
	encounterRepo := repo.NewEncounterRepository(db)
	return &EncounterService{EncounterRepo: encounterRepo}
}

func (service *EncounterService) Create(encounter *model.Encounter) error {
	err := service.EncounterRepo.CreateEncounter(encounter)
	if err != nil {
		return err
	}
	return nil
}

func (service *EncounterService) Find(id int) (*model.Encounter, error) {
	tour, err := service.EncounterRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("execution with id %d not found", id)
	}
	return &tour, nil
}

func (service *EncounterService) FindByCheckPointId(id int) (*model.Encounter, error) {
	tour, err := service.EncounterRepo.FindByCheckPointId(id)
	if err != nil {
		return nil, fmt.Errorf("execution with id %d not found", id)
	}
	return &tour, nil
}

func (service *EncounterService) Delete(id int) error {
	err := service.EncounterRepo.DeleteById(id)

	if err != nil {
		return err
	}
	return nil
}


/*

func (service *EncounterService) CreateForTourist(encounter *model.Encounter, checkpointId int64, isSecretPrerequisite bool, userId int64) (*model.Encounter, error) {

}

func (service *EncounterService) getTouristLevel(userId int64) (int, error) {
	// Implementirajte logiku za dobijanje nivoa turiste iz baze podataka ili drugog izvora
	// Ova funkcija treba da vrati nivo turiste i eventualnu grešku
	return 0, nil
}
// Funkcija za mapiranje EncounterDto u odgovarajući Encounter tip
func MapEncounter(service *EncounterService) (encounter *model.Encounter) {
	switch encounter.Type {
	case "Location":
		return service.SocialEncounterService.CreateSocialEncounter(encounter)
	case "Social":
		return &SocialEncounter{}, nil
	default:
		return nil, errors.New("Unknown encounter type")
	}
}

func (service *EncounterService) setEncounterOnCheckpoint(checkpointId int64, encounterId int64, isSecretPrerequisite bool, authorId int64) error {
	// Implementirajte logiku za postavljanje susreta na kontrolnoj tački
	// Ova funkcija treba da vrati eventualnu grešku
	return nil
}*/
  
  func (service *EncounterService) Update(encounter *model.Encounter) error {
	err := service.EncounterRepo.Update(encounter)
    if err != nil {
		return err
	}
	return nil
}

