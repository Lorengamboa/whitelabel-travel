package data

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// GetAll returns all users from the database
func (um CustomerModel) GetAll() ([]*Customer, error) {
	query := `
		SELECT 
			c.id, c.email, c.name, c.phone_number, c.Address, c.logo, c.url, c.date_joined
		FROM 
			customers c
	`
	var customers []*Customer
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := um.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Email, &customer.Name, &customer.PhoneNumber, &customer.Address, &customer.Logo, &customer.URL, &customer.DateJoined)
		if err != nil {
			return nil, err
		}

		customers = append(customers, &customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

// Get returns a single user from the database
func (um CustomerModel) Get(id *uuid.UUID) (*Customer, error) {
	query := `
		SELECT 
			c.id, c.email, c.name, c.phone_number, c.Address, c.logo, c.url, c.date_joined
		FROM 
			customers c
		WHERE 
			c.id = $1
	`
	var customer Customer
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := um.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&customer.ID, &customer.Email, &customer.Name, &customer.PhoneNumber, &customer.Address, &customer.Logo, &customer.URL, &customer.DateJoined)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
