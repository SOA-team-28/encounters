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
