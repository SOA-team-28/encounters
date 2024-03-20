package service

import (
	"database-example/repo"
)

type SocialEncounterService struct {
	EncounterRepo *repo.SocialEncounterRepository
}
