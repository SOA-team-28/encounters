package service

import (
	"database-example/model"
	"database-example/repo"

	"gorm.io/gorm"
)

type EncounterExecutionService struct {
	EncounterExecutionRepo *repo.EncounterExecutionRepository
}

func NewEncounterExecutionService(db *gorm.DB) *EncounterExecutionService {
	encounterExecutionRepo := repo.NewEncounterExecutionRepository(db)
	return &EncounterExecutionService{EncounterExecutionRepo: encounterExecutionRepo}
}

func (service *EncounterExecutionService) Create(encounter *model.EncounterExecution) error {
	err := service.EncounterExecutionRepo.CreateEncounterExecution(encounter)
	if err != nil {
		return err
	}
	return nil
}

func (service *EncounterExecutionService) GetAllCompletedByTourist(touristID int64) ([]model.EncounterExecution, error) {
	encounterExecutions, err := service.EncounterExecutionRepo.GetAllCompletedByTourist(touristID)
	if err != nil {
		return nil, err
	}
	return encounterExecutions, nil
}
