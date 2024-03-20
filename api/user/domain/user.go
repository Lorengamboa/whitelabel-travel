package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type RoleNames string

type Profile struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	PhoneNumber string    `json:"phone_number"`
	BirthDate   time.Time `json:"birth_date"`
}

type User struct {
	ID         uuid.UUID `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"_"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	IsActive   bool      `json:"is_active"`
	Role       RoleNames `json:"role"`
	Thumbnail  string    `json:"thumbnail"`
	DateJoined time.Time `json:"date_joined"`
	Profile    *Profile  `json:"profile"`
}

const (
	Superadmin RoleNames = "super_admin"
)

type UserRepository interface {
	GetById(ctx context.Context, id uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetAll(ctx context.Context) ([]User, error)
	Insert(ctx context.Context, user *User) error
}
type UserUsecase interface {
	GetById(ctx context.Context, id uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetAll(ctx context.Context) ([]User, error)
	Insert(ctx context.Context, user *User) error
}
