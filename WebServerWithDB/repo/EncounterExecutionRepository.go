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

func (repo *EncounterExecutionRepository) DeleteById(id int) error {
	// Prvo pronađite susret koji želite obrisati
	var encounterExecution model.EncounterExecution
	if err := repo.DatabaseConnection.First(&encounterExecution, id).Error; err != nil {
		return err
	}

	// Zatim obrišite susret iz baze podataka
	if err := repo.DatabaseConnection.Delete(&encounterExecution).Error; err != nil {
		return err
	}

	return nil
}
