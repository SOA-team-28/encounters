package model

/*
// EncounterStatus enumeracija
type EncounterStatus int

const (
	Draft EncounterStatus = iota
	Archived
	Published
)

// EncounterType enumeracija
type EncounterType int

const (
	Social EncounterType = iota
	Location
	Misc
)
*/

type Encounter struct {
	AuthorId    int
	Description string
	Id          int
	Latitude    int
	Longitude   int
	Name        string
	Status      string
	Type        string
	XP          int
}
