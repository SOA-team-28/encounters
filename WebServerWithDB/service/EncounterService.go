package service

import (
	"database-example/model"
	"database-example/repo"
)

type EncounterService struct {
	EncounterRepo *repo.EncounterRepository
}

func (service *EncounterService) Create(encounter *model.Encounter) error {
	err := service.EncounterRepo.CreateEncounter(encounter)
	if err != nil {
		return err
	}
	return nil
}
