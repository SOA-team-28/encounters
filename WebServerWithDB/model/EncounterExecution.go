package model

import (
	"time"
)

type EncounterExecution struct {
	Id          int
	EncounterId int
	TouristId   int
	Status      string
	StartTime   *time.Time
	Encounter   Encounter
}
