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

func (repo *EncounterExecutionRepository) UpdateStatusByCheckPointId(checkPointId int) error {
	var encounterExecution model.EncounterExecution

	dbResult := repo.DatabaseConnection.
		Joins("JOIN encounters ON encounter_executions.encounter_id = encounters.id").
		Where("encounters.check_point_id = ?", checkPointId).
		Find(&encounterExecution)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	encounterExecution.Status = "Completed"

	dbResult = repo.DatabaseConnection.Save(&encounterExecution)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	println("Updated successfully! ")
	return nil
}
