package service

import (
	"database-example/repo"
)

type HiddenLocationEncounterService struct {
	EncounterRepo *repo.HiddenLocationEncounterRepository
}
