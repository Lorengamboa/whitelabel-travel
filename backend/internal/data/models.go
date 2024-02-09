package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("a user with these details was not found")
)

type Models struct {
	Users     UserModel
	Customers CustomerModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:     UserModel{DB: db},
		Customers: CustomerModel{DB: db},
	}
}
