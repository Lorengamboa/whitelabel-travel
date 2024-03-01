package data

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// get clients
func (um ClientModel) GetClients(id *uuid.UUID) ([]*Client, error) {
	query := `
		SELECT 
			id, title, url, logo, primary_color, secondary_color, name, email, address1, address2, country, city, date_joined, customer_id
		FROM 
			clients
		WHERE 
			customer_id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := um.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	clients := []*Client{}
	for rows.Next() {
		client := &Client{}
		err := rows.Scan(&client.ID, &client.Title, &client.URL, &client.Logo, &client.PrimaryColor, &client.SecondaryColor, &client.Name, &client.Email, &client.Address, &client.Address2, &client.Country, &client.City, &client.DateJoined, &client.CustomerID)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}

// Inserts a new client that belongs to an existing customer
func (um ClientModel) InsertClient(client *Client) (*uuid.UUID, error) {
	query := `
		INSERT INTO 
			clients (title, url, logo, primary_color, secondary_color, name, email, address1, address2, country, city, date_joined, customer_id)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := um.DB.QueryRowContext(ctx, query, client.Title, client.URL, client.Logo, client.PrimaryColor, client.SecondaryColor, client.Name, client.Email, client.Address, client.Address2, client.Country, client.City, client.DateJoined, client.CustomerID).Scan(&client.ID)
	if err != nil {
		return nil, err
	}

	return &client.ID, nil
}

// Get a single client
func (um ClientModel) GetClient(customerID *uuid.UUID, clientID string) (*Client, error) {
	query := `
		SELECT 
			id, title, url, logo, primary_color, secondary_color, name, email, address1, address2, country, city, date_joined, customer_id
		FROM 
			clients
		WHERE 
			customer_id = $1 AND id = $2
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := &Client{}
	err := um.DB.QueryRowContext(ctx, query, customerID, clientID).Scan(&client.ID, &client.Title, &client.URL, &client.Logo, &client.PrimaryColor, &client.SecondaryColor, &client.Name, &client.Email, &client.Address, &client.Address2, &client.Country, &client.City, &client.DateJoined, &client.CustomerID)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Get all travel packages for a client
func (um TravelPackagesModel) GetTravelPackages(customerID *uuid.UUID, clientID string) ([]*TravelPackages, error) {
	query := `
		SELECT 
			id, client_id, package_name, duration, itinerary, package_includes, package_excludes, recommended_gear, difficulty_level, price
		FROM 
			travel_packages
		WHERE 
			client_id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := um.DB.QueryContext(ctx, query, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	packages := []*TravelPackages{}
	for rows.Next() {
		pkg := &TravelPackages{}
		err := rows.Scan(&pkg.ID, &pkg.ClientID, &pkg.PackageName, &pkg.Duration, &pkg.Itinerary, &pkg.PackageIncludes, &pkg.PackageExcludes, &pkg.RecommendedGear, &pkg.DifficultyLevel, &pkg.Price)
		if err != nil {
			return nil, err
		}
		packages = append(packages, pkg)
	}

	return packages, nil
}

// Insert a new travel package for a client
func (um TravelPackagesModel) InsertTravelPackage(customerID *uuid.UUID, clientID string, pkg *TravelPackages) (*uuid.UUID, error) {
	query := `
		INSERT INTO 
			travel_packages (client_id, package_name, duration, itinerary, package_includes, package_excludes, recommended_gear, difficulty_level, price)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	travelPackage := &TravelPackages{}

	err := um.DB.QueryRowContext(ctx, query, clientID, pkg.PackageName, pkg.Duration, pkg.Itinerary, pkg.PackageIncludes, pkg.PackageExcludes, pkg.RecommendedGear, pkg.DifficultyLevel, pkg.Price).Scan(&clientID)
	if err != nil {
		return nil, err
	}

	return travelPackage.ID, nil
}

// Get a single travel package for a client
func (um TravelPackagesModel) GetTravelPackage(customerID *uuid.UUID, clientID string, travelPackageID string) (*TravelPackages, error) {
	query := `
		SELECT 
			id, client_id, package_name, duration, itinerary, package_includes, package_excludes, recommended_gear, difficulty_level, price
		FROM 
			travel_packages
		WHERE 
			client_id = $1 AND id = $2
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	travelPackage := &TravelPackages{}
	err := um.DB.QueryRowContext(ctx, query, clientID, travelPackageID).Scan(&travelPackage.ID, &travelPackage.ClientID, &travelPackage.PackageName, &travelPackage.Duration, &travelPackage.Itinerary, &travelPackage.PackageIncludes, &travelPackage.PackageExcludes, &travelPackage.RecommendedGear, &travelPackage.DifficultyLevel, &travelPackage.Price)
	if err != nil {
		return nil, err
	}

	return travelPackage, nil
}

// Public API's
func (um ClientModel) GetPublicClient(clientID string) (*Client, error) {
	query := `
		SELECT 
			id, title, url, logo, primary_color, secondary_color, name, email, address1, address2, country, city, date_joined, customer_id
		FROM 
			clients
		WHERE 
			id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := &Client{}
	err := um.DB.QueryRowContext(ctx, query, clientID).Scan(&client.ID, &client.Title, &client.URL, &client.Logo, &client.PrimaryColor, &client.SecondaryColor, &client.Name, &client.Email, &client.Address, &client.Address2, &client.Country, &client.City, &client.DateJoined, &client.CustomerID)
	if err != nil {
		return nil, err
	}

	return client, nil
}
