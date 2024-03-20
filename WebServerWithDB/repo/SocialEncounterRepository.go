package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type SocialEncounterRepository struct {
	DatabaseConnection *gorm.DB
}

func NewSocialEncounterRepository(databaseConnection *gorm.DB) *SocialEncounterRepository {
	return &SocialEncounterRepository{DatabaseConnection: databaseConnection}
}

func (repo *HiddenLocationEncounterRepository) CreateSocialEncounter(encounter *model.SocialEncounter) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
