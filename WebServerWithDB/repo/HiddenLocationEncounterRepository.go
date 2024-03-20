package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type HiddenLocationEncounterRepository struct {
	DatabaseConnection *gorm.DB
}

func NewHiddenLocationEncounterRepository(databaseConnection *gorm.DB) *HiddenLocationEncounterRepository {
	return &HiddenLocationEncounterRepository{DatabaseConnection: databaseConnection}
}

func (repo *HiddenLocationEncounterRepository) CreateHiddenLocationEncounter(encounter *model.HiddenLocationEncounter) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
