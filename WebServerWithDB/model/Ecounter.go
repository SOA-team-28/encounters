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
	AuthorId          int
	Description       string
	Id                int
	Latitude          float64
	Longitude         float64
	Name              string
	Status            string
	Type              string
	XP                int
	LocationLongitude float64
	LocationLatitude  float64
	Image             string
	Range             float64
	RequiredPeople    int
	CheckPointId      int
}
