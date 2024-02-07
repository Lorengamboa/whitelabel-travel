package data

import (
	"database/sql"
	"errors"
	"time"
	"v0/internal/types"

	"github.com/google/uuid"
)

type RoleNames string

const (
	Superadmin RoleNames = "super_admin"
)

type UserProfile struct {
	ID          *uuid.UUID     `json:"id"`
	UserID      *uuid.UUID     `json:"user_id"`
	PhoneNumber *string        `json:"phone_number"`
	BirthDate   types.NullTime `json:"birth_date"`
}

type User struct {
	ID         uuid.UUID   `json:"id"`
	Email      string      `json:"email"`
	Password   password    `json:"_"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	IsActive   bool        `json:"is_active"`
	Role       RoleNames   `json:"role"`
	Thumbnail  *string     `json:"thumbnail"`
	DateJoined time.Time   `json:"date_joined"`
	Profile    UserProfile `json:"profile"`
}

type password struct {
	plaintext *string
	hash      string
}

type UserModel struct {
	DB *sql.DB
}

type UserID struct {
	Id uuid.UUID
}

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)
