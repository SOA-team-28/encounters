package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type EncounterRequestRepository struct {
	DatabaseConnection *gorm.DB
}

func NewEncounterRequestRepository(databaseConnection *gorm.DB) *EncounterRequestRepository {
	return &EncounterRequestRepository{DatabaseConnection: databaseConnection}
}

func (repo *EncounterRequestRepository) FindAll() ([]model.EncounterRequest, error) {
	var requests []model.EncounterRequest
	dbResult := repo.DatabaseConnection.Find(&requests)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return requests, nil
}

func (repo *EncounterRequestRepository) Accept(id int) error {
	encounterRequest := model.EncounterRequest{}
	dbResult := repo.DatabaseConnection.First(&encounterRequest, "id = ?", id)

	if dbResult.Error != nil {
		return dbResult.Error
	}

	encounterRequest.Status = "Accepted"

	dbResult = repo.DatabaseConnection.Save(&encounterRequest)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	println("Updated successfully! ")
	return nil
}
func (repo *EncounterRequestRepository) Reject(id int) error {
	encounterRequest := model.EncounterRequest{}
	dbResult := repo.DatabaseConnection.First(&encounterRequest, "id = ?", id)

	if dbResult.Error != nil {
		return dbResult.Error
	}

	encounterRequest.Status = "Rejected"

	dbResult = repo.DatabaseConnection.Save(&encounterRequest)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	println("Updated successfully! ")
	return nil
}
