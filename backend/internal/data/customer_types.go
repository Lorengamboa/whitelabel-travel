package data

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID          uuid.UUID       `json:"id"`
	Email       string          `json:"email"`
	Name        string          `json:"name"`
	PhoneNumber *string         `json:"phone_number"`
	Address     string          `json:"address"`
	Logo        *sql.NullString `json:"logo"`
	URL         string          `json:"url"`
	DateJoined  time.Time       `json:"date_joined"`
}

type CustomerModel struct {
	DB *sql.DB
}

type CustomerID struct {
	Id uuid.UUID
}
