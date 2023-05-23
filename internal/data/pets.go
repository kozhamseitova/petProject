package data

import (
	"time"
)

type Pet struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
	SpeciesID int64     `json:"species_id"`
}
