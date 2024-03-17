package service

import (
	"database-example/model"
	"database-example/repo"

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
