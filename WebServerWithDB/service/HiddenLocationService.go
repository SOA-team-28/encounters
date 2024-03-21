package service

import (
	"database-example/model"
	"database-example/repo"

	"gorm.io/gorm"
)

type HiddenLocationEncounterService struct {
	HiddenLocationEncounterRepository *repo.HiddenLocationEncounterRepository
}

func NewHiddenLocationEncounterService(db *gorm.DB) *HiddenLocationEncounterService {
	hiddenLocationEncounterRepo := repo.NewHiddenLocationEncounterRepository(db)
	return &HiddenLocationEncounterService{HiddenLocationEncounterRepository: hiddenLocationEncounterRepo}
}

func (service *HiddenLocationEncounterService) CreateHiddenLocationEncounter(encounter *model.Encounter) error {
	var hiddenLocaionEncounter model.HiddenLocationEncounter
	hiddenLocaionEncounter.Image = encounter.Image
	hiddenLocaionEncounter.LocationLatitude = encounter.LocationLatitude
	hiddenLocaionEncounter.LocationLongitude = encounter.LocationLongitude
	hiddenLocaionEncounter.Range = encounter.Range
	err := service.HiddenLocationEncounterRepository.CreateHiddenLocationEncounterByTourist(&hiddenLocaionEncounter)
	if err != nil {
		return err
	}
	return nil
}
