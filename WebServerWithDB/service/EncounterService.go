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

func (service *EncounterService) Update(encounter *model.Encounter) error {
	err := service.EncounterRepo.Update(encounter)
	if err != nil {
		return err
	}
	return nil
}
