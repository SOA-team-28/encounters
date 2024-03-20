package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"

	"gorm.io/gorm"
)

type EncounterRequestService struct {
	EncounterRequestRepo *repo.EncounterRequestRepository
}

func NewEncounterRequestService(db *gorm.DB) *EncounterRequestService {
	encounterRequestRepo := repo.NewEncounterRequestRepository(db)
	return &EncounterRequestService{EncounterRequestRepo: encounterRequestRepo}
}
func (service *EncounterRequestService) FindAll() ([]model.EncounterRequest, error) {
	requests, err := service.EncounterRequestRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("requests not found: %v", err)
	}
	return requests, nil
}

func (service *EncounterRequestService) Accept(id int) error {
	err := service.EncounterRequestRepo.Accept(id)
	if err != nil {
		return err
	}
	return nil
}
func (service *EncounterRequestService) Reject(id int) error {
	err := service.EncounterRequestRepo.Reject(id)
	if err != nil {
		return err
	}
	return nil
}
