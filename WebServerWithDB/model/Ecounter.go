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
	AuthorId    int    `json:"authorId"`
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	XP          int    `json:"xp"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Longitude   int    `json:"longitude"`
	Latitude    int    `json:"latitude"`
}
