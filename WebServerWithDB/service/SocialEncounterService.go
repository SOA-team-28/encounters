package service

import (
	"database-example/model"
	"database-example/repo"

	"gorm.io/gorm"
)

type SocialEncounterService struct {
	SocialEncounterRepo *repo.SocialEncounterRepository
}

func NewSocialEncounterService(db *gorm.DB) *SocialEncounterService {
	hiddenLocationEncounterRepo := repo.NewSocialEncounterRepository(db)
	return &SocialEncounterService{SocialEncounterRepo: hiddenLocationEncounterRepo}
}

func (service *SocialEncounterService) CreateSocialEncounterByTourist(encounter *model.Encounter) error {
	var socialEncounter model.SocialEncounter
	socialEncounter.Range = encounter.Range
	socialEncounter.RequiredPeople = encounter.RequiredPeople
	err := service.SocialEncounterRepo.CreateSocialEncounterByTourist(&socialEncounter)
	if err != nil {
		return err
	}
	return nil
}
