package main

import (
	"errors"
	"net/http"
	"v0/cmd/deploy"
	"v0/internal/data"

	"github.com/google/uuid"
)

func (app *application) getCustomersHandler(w http.ResponseWriter, r *http.Request) {
	userID, status, err := app.extractParamsFromSession(r)
	if err != nil {
		switch *status {
		case http.StatusUnauthorized:
			app.unauthorizedResponse(w, r, err)

		case http.StatusBadRequest:
			app.badRequestResponse(w, r, errors.New("invalid cookie"))

		case http.StatusInternalServerError:
			app.serverErrorResponse(w, r, err)

		default:
			app.serverErrorResponse(
				w,
				r,
				errors.New("something happened and we could not fullfil your request at the moment"),
			)
		}
		return
	}

	db_user, err := app.models.Users.Get(userID.Id)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if db_user.Role == "RoleAdmin" {
		app.unauthorizedResponse(w, r, errors.New("you are not authorized to access this resource"))
		return
	}

	db_customers, err := app.models.Customers.GetAll()

	app.writeJSON(w, http.StatusOK, db_customers, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	app.logSuccess(r, http.StatusOK, db_customers)
}

// getCustomerHandler /customers/{id}
func (app *application) getCustomerHandler(w http.ResponseWriter, r *http.Request) {
	userID, status, err := app.extractParamsFromSession(r)
	if err != nil {
		switch *status {
		case http.StatusUnauthorized:
			app.unauthorizedResponse(w, r, err)

		case http.StatusBadRequest:
			app.badRequestResponse(w, r, errors.New("invalid cookie"))

		case http.StatusInternalServerError:
			app.serverErrorResponse(w, r, err)

		default:
			app.serverErrorResponse(
				w,
				r,
				errors.New("something happened and we could not fullfil your request at the moment"),
			)
		}
		return
	}

	db_user, err := app.models.Users.Get(userID.Id)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if db_user.Role == "RoleAdmin" {
		app.unauthorizedResponse(w, r, errors.New("you are not authorized to access this resource"))
		return
	}

	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	db_customer, err := app.models.Customers.Get(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, db_customer, nil)
	app.logSuccess(r, http.StatusOK, db_customer)
}

// deleteCustomerHandler /customers/{id}
func (app *application) deleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	userID, status, err := app.extractParamsFromSession(r)
	if err != nil {
		switch *status {
		case http.StatusUnauthorized:
			app.unauthorizedResponse(w, r, err)

		case http.StatusBadRequest:
			app.badRequestResponse(w, r, errors.New("invalid cookie"))

		case http.StatusInternalServerError:
			app.serverErrorResponse(w, r, err)

		default:
			app.serverErrorResponse(
				w,
				r,
				errors.New("something happened and we could not fullfil your request at the moment"),
			)
		}
		return
	}

	db_user, err := app.models.Users.Get(userID.Id)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if db_user.Role == "RoleAdmin" {
		app.unauthorizedResponse(w, r, errors.New("you are not authorized to access this resource"))
		return
	}

	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.models.Customers.Delete(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, "customer deleted successfully", nil)
	app.logSuccess(r, http.StatusOK, "customer deleted successfully")
}

// createCustomerHandler /customers
func (app *application) createCustomerHandler(w http.ResponseWriter, r *http.Request) {
	userID, status, err := app.extractParamsFromSession(r)
	if err != nil {
		switch *status {
		case http.StatusUnauthorized:
			app.unauthorizedResponse(w, r, err)

		case http.StatusBadRequest:
			app.badRequestResponse(w, r, errors.New("invalid cookie"))

		case http.StatusInternalServerError:
			app.serverErrorResponse(w, r, err)

		default:
			app.serverErrorResponse(
				w,
				r,
				errors.New("something happened and we could not fullfil your request at the moment"),
			)
		}
		return
	}

	db_user, err := app.models.Users.Get(userID.Id)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if db_user.Role == "RoleAdmin" {
		app.unauthorizedResponse(w, r, errors.New("you are not authorized to access this resource"))
		return
	}

	var input struct {
		Logo        string  `json:"logo"`
		Name        string  `json:"name"`
		Email       string  `json:"email"`
		Address     string  `json:"address"`
		PhoneNumber *string `json:"phone_number"`
		URL         string  `json:"url"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	db_customer := &data.Customer{
		Logo:        input.Logo,
		Name:        input.Name,
		Email:       input.Email,
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
		URL:         input.URL,
	}

	id, err := app.models.Customers.Insert(db_customer)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, id, nil)
	app.logSuccess(r, http.StatusCreated, id)
}

// getCustomerClientsHandler /customers/{id}/clients
func (app *application) getCustomerClientsHandler(w http.ResponseWriter, r *http.Request) {
	userID, status, err := app.extractParamsFromSession(r)
	if err != nil {
		switch *status {
		case http.StatusUnauthorized:
			app.unauthorizedResponse(w, r, err)

		case http.StatusBadRequest:
			app.badRequestResponse(w, r, errors.New("invalid cookie"))

		case http.StatusInternalServerError:
			app.serverErrorResponse(w, r, err)

		default:
			app.serverErrorResponse(
				w,
				r,
				errors.New("something happened and we could not fullfil your request at the moment"),
			)
		}
		return
	}

	db_user, err := app.models.Users.Get(userID.Id)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if db_user.Role == "RoleAdmin" {
		app.unauthorizedResponse(w, r, errors.New("you are not authorized to access this resource"))
		return
	}

	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	db_clients, err := app.models.Clients.GetClients(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, db_clients, nil)
	app.logSuccess(r, http.StatusOK, db_clients)
}

// createClientHandler /customers/{id}/clients
func (app *application) createClientHandler(w http.ResponseWriter, r *http.Request) {
	userID, status, err := app.extractParamsFromSession(r)
	if err != nil {
		switch *status {
		case http.StatusUnauthorized:
			app.unauthorizedResponse(w, r, err)

		case http.StatusBadRequest:
			app.badRequestResponse(w, r, errors.New("invalid cookie"))

		case http.StatusInternalServerError:
			app.serverErrorResponse(w, r, err)

		default:
			app.serverErrorResponse(
				w,
				r,
				errors.New("something happened and we could not fullfil your request at the moment"),
			)
		}
		return
	}

	db_user, err := app.models.Users.Get(userID.Id)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if db_user.Role == "RoleAdmin" {
		app.unauthorizedResponse(w, r, errors.New("you are not authorized to access this resource"))
		return
	}

	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var input struct {
		CustomerId     *uuid.UUID `json:"customerId"`
		URL            string     `json:"url"`
		Title          string     `json:"title"`
		PrimaryColor   string     `json:"primaryColor"`
		SecondaryColor string     `json:"secondaryColor"`
		Logo           string     `json:"logo"`
		Address        string     `json:"address1"`
		Address2       string     `json:"address2"`
		Country        string     `json:"country"`
		City           string     `json:"city"`
		Email          string     `json:"email"`
		Name           string     `json:"name"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	db_client := &data.Client{
		CustomerID:     id,
		URL:            input.URL,
		Title:          input.Title,
		PrimaryColor:   input.PrimaryColor,
		SecondaryColor: input.SecondaryColor,
		Logo:           input.Logo,
		Address:        input.Address,
		Address2:       input.Address2,
		Country:        input.Country,
		City:           input.City,
		Email:          input.Email,
		Name:           input.Name,
	}

	id, err = app.models.Clients.InsertClient(db_client)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, id, nil)
	app.logSuccess(r, http.StatusCreated, id)
}

// getClientHandler /customers/{id}/clients/{id}
func (app *application) getClientHandler(w http.ResponseWriter, r *http.Request) {
	userID, status, err := app.extractParamsFromSession(r)
	if err != nil {
		switch *status {
		case http.StatusUnauthorized:
			app.unauthorizedResponse(w, r, err)

		case http.StatusBadRequest:
			app.badRequestResponse(w, r, errors.New("invalid cookie"))

		case http.StatusInternalServerError:
			app.serverErrorResponse(w, r, err)

		default:
			app.serverErrorResponse(
				w,
				r,
				errors.New("something happened and we could not fullfil your request at the moment"),
			)
		}
		return
	}

	db_user, err := app.models.Users.Get(userID.Id)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if db_user.Role == "RoleAdmin" {
		app.unauthorizedResponse(w, r, errors.New("you are not authorized to access this resource"))
		return
	}

	// read both params
	customerID, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	clientID, err := app.getParam(r, "clientId")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	db_client, err := app.models.Clients.GetClient(customerID, clientID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, db_client, nil)
	app.logSuccess(r, http.StatusOK, db_client)
}

// deployClientHandler /customers/{id}/clients/{id}/deploy
func (app *application) deployClientHandler(w http.ResponseWriter, r *http.Request) {
	userID, status, err := app.extractParamsFromSession(r)
	if err != nil {
		switch *status {
		case http.StatusUnauthorized:
			app.unauthorizedResponse(w, r, err)

		case http.StatusBadRequest:
			app.badRequestResponse(w, r, errors.New("invalid cookie"))

		case http.StatusInternalServerError:
			app.serverErrorResponse(w, r, err)

		default:
			app.serverErrorResponse(
				w,
				r,
				errors.New("something happened and we could not fullfil your request at the moment"),
			)
		}
		return
	}

	db_user, err := app.models.Users.Get(userID.Id)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if db_user.Role == "RoleAdmin" {
		app.unauthorizedResponse(w, r, errors.New("you are not authorized to access this resource"))
		return
	}

	// read both params
	customerID, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	clientID, err := app.getParam(r, "clientId")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	db_client, err := app.models.Clients.GetClient(customerID, clientID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// deploy client
	err = deploy.DeployClient(db_client)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, db_client, nil)
	app.logSuccess(r, http.StatusOK, db_client)
}

// getClientTravelPackagesHandler /customers/{id}/clients/{id}/travel-packages
func (app *application) getClientTravelPackagesHandler(w http.ResponseWriter, r *http.Request) {
	userID, status, err := app.extractParamsFromSession(r)
	if err != nil {
		switch *status {
		case http.StatusUnauthorized:
			app.unauthorizedResponse(w, r, err)

		case http.StatusBadRequest:
			app.badRequestResponse(w, r, errors.New("invalid cookie"))

		case http.StatusInternalServerError:
			app.serverErrorResponse(w, r, err)

		default:
			app.serverErrorResponse(
				w,
				r,
				errors.New("something happened and we could not fullfil your request at the moment"),
			)
		}
		return
	}

	db_user, err := app.models.Users.Get(userID.Id)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if db_user.Role == "RoleAdmin" {
		app.unauthorizedResponse(w, r, errors.New("you are not authorized to access this resource"))
		return
	}

	// read both params
	customerID, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	clientID, err := app.getParam(r, "clientId")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	db_travelPackages, err := app.models.TravelPackages.GetTravelPackages(customerID, clientID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, db_travelPackages, nil)
	app.logSuccess(r, http.StatusOK, db_travelPackages)
}

// createClientTravelPackageHandler /customers/{id}/clients/{id}/travel-packages
func (app *application) createClientTravelPackageHandler(w http.ResponseWriter, r *http.Request) {
	userID, status, err := app.extractParamsFromSession(r)
	if err != nil {
		switch *status {
		case http.StatusUnauthorized:
			app.unauthorizedResponse(w, r, err)

		case http.StatusBadRequest:
			app.badRequestResponse(w, r, errors.New("invalid cookie"))

		case http.StatusInternalServerError:
			app.serverErrorResponse(w, r, err)

		default:
			app.serverErrorResponse(
				w,
				r,
				errors.New("something happened and we could not fullfil your request at the moment"),
			)
		}
		return
	}

	db_user, err := app.models.Users.Get(userID.Id)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if db_user.Role == "RoleAdmin" {
		app.unauthorizedResponse(w, r, errors.New("you are not authorized to access this resource"))
		return
	}

	// read both params
	customerID, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	clientID, err := app.getParam(r, "clientId")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var input struct {
		PackageName     string  `json:"packageName"`
		Duration        int     `json:"duration"`
		Itinerary       string  `json:"itinerary"`
		PackageIncludes string  `json:"packageIncludes"`
		PackageExcludes string  `json:"packageExcludes"`
		RecommendedGear string  `json:"recommendedGear"`
		DifficultyLevel string  `json:"difficultyLevel"`
		Price           float64 `json:"price"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	db_travelPackage := &data.TravelPackages{
		PackageName:     input.PackageName,
		Duration:        input.Duration,
		Itinerary:       input.Itinerary,
		PackageIncludes: input.PackageIncludes,
		PackageExcludes: input.PackageExcludes,
		RecommendedGear: input.RecommendedGear,
		DifficultyLevel: input.DifficultyLevel,
		Price:           float64(input.Price),
	}

	id, err := app.models.TravelPackages.InsertTravelPackage(customerID, clientID, db_travelPackage)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, id, nil)
	app.logSuccess(r, http.StatusCreated, id)
}

// getClientTravelPackageHandler /customers/{id}/clients/{id}/travel-packages/{travelPackageId}
func (app *application) getClientTravelPackageHandler(w http.ResponseWriter, r *http.Request) {
	userID, status, err := app.extractParamsFromSession(r)
	if err != nil {
		switch *status {
		case http.StatusUnauthorized:
			app.unauthorizedResponse(w, r, err)

		case http.StatusBadRequest:
			app.badRequestResponse(w, r, errors.New("invalid cookie"))

		case http.StatusInternalServerError:
			app.serverErrorResponse(w, r, err)

		default:
			app.serverErrorResponse(
				w,
				r,
				errors.New("something happened and we could not fullfil your request at the moment"),
			)
		}
		return
	}

	db_user, err := app.models.Users.Get(userID.Id)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if db_user.Role == "RoleAdmin" {
		app.unauthorizedResponse(w, r, errors.New("you are not authorized to access this resource"))
		return
	}

	// read both params
	customerID, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	clientID, err := app.getParam(r, "clientId")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	travelPackageID, err := app.getParam(r, "travelPackageId")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	db_travelPackage, err := app.models.TravelPackages.GetTravelPackage(customerID, clientID, travelPackageID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, db_travelPackage, nil)
	app.logSuccess(r, http.StatusOK, db_travelPackage)
}

// getPublicClientHandler /public/client/{id}
func (app *application) getPublicClientHandler(w http.ResponseWriter, r *http.Request) {
	clientID, err := app.getParam(r, "clientId")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	db_client, err := app.models.Clients.GetPublicClient(clientID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, db_client, nil)
	app.logSuccess(r, http.StatusOK, db_client)
}
