package data

import (
	"database/sql"

	"github.com/google/uuid"
)

type TravelPackages struct {
	ID              *uuid.UUID `json:"id"`
	ClientID        uuid.UUID  `json:"clientId"`
	PackageName     string     `json:"packageName"`
	Duration        int        `json:"duration"`
	Itinerary       string     `json:"itinerary"`
	PackageIncludes string     `json:"packageIncludes"`
	PackageExcludes string     `json:"packageExcludes"`
	RecommendedGear string     `json:"recommendedGear"`
	DifficultyLevel string     `json:"difficultyLevel"`
	Price           float64    `json:"price"`
}

type TravelPackagesModel struct {
	DB *sql.DB
}

type TravelPackagesID struct {
	Id uuid.UUID
}
