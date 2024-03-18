package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type EncounterExecutionRepository struct {
	DatabaseConnection *gorm.DB
}

func NewEncounterExecutionRepository(databaseConnection *gorm.DB) *EncounterExecutionRepository {
	return &EncounterExecutionRepository{DatabaseConnection: databaseConnection}
}

func (repo *EncounterExecutionRepository) CreateEncounterExecution(encounterExecution *model.EncounterExecution) error {
	dbResult := repo.DatabaseConnection.Create(encounterExecution)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterExecutionRepository) GetAllCompletedByTourist(touristID int64) ([]model.EncounterExecution, error) {
	var encounterExecutions []model.EncounterExecution

	// Izvršite upit koristeći GORM za dohvat završenih susreta za turistu
	result := repo.DatabaseConnection.
		Preload("Encounter").
		Where("tourist_id = ? AND status = ?", touristID, "Completed").
		Find(&encounterExecutions)

	if result.Error != nil {
		return nil, result.Error
	}

	return encounterExecutions, nil
}
