package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type EncounterRepository struct {
	DatabaseConnection *gorm.DB
}

func NewEncounterRepository(databaseConnection *gorm.DB) *EncounterRepository {
	return &EncounterRepository{DatabaseConnection: databaseConnection}
}

func (repo *EncounterRepository) CreateEncounter(encounter *model.Encounter) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterRepository) FindById(id int) (model.Encounter, error) {
	encounter := model.Encounter{}
	dbResult := repo.DatabaseConnection.First(&encounter, "id = ?", id)
	if dbResult.Error != nil {
		return encounter, dbResult.Error
	}
	return encounter, nil
}

func (repo *EncounterRepository) FindByCheckPointId(id int) (model.Encounter, error) {
	encounter := model.Encounter{}
	dbResult := repo.DatabaseConnection.First(&encounter, "check_point_id = ?", id)
	if dbResult.Error != nil {
		return encounter, dbResult.Error
	}
	return encounter, nil
}
