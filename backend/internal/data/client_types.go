package data

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID             uuid.UUID  `json:"id"`
	CustomerID     *uuid.UUID `json:"customerId"`
	URL            string     `json:"url"`
	Title          string     `json:"title"`
	PrimaryColor   string     `json:"primary_color"`
	SecondaryColor string     `json:"secondary_color"`
	Logo           string     `json:"logo"`
	Address        string     `json:"address1"`
	Address2       string     `json:"address2"`
	City           string     `json:"city"`
	Country        string     `json:"country"`
	Email          string     `json:"email"`
	Name           string     `json:"name"`
	DateJoined     time.Time  `json:"date_joined"`
}

type ClientModel struct {
	DB *sql.DB
}

type ClientID struct {
	Id uuid.UUID
}
