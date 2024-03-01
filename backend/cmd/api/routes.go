package main

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// User-related routes
	router.HandlerFunc(http.MethodPost, "/users/register", app.registerUserHandler)
	router.HandlerFunc(http.MethodPost, "/users/login", app.loginUserHandler)
	router.HandlerFunc(http.MethodGet, "/users", app.getUsersHandler)

	// Customer-related routes
	router.HandlerFunc(http.MethodGet, "/customers", app.getCustomersHandler)
	router.HandlerFunc(http.MethodPost, "/customers", app.createCustomerHandler)
	router.HandlerFunc(http.MethodGet, "/customers/:id", app.getCustomerHandler)
	router.HandlerFunc(http.MethodDelete, "/customers/:id", app.deleteCustomerHandler)

	// Client-related routes
	router.HandlerFunc(http.MethodGet, "/customers/:id/clients", app.getCustomerClientsHandler)
	router.HandlerFunc(http.MethodPost, "/customers/:id/clients", app.createClientHandler)
	router.HandlerFunc(http.MethodGet, "/customers/:id/clients/:clientId", app.getClientHandler)
	router.HandlerFunc(http.MethodGet, "/customers/:id/clients/:clientId/travel-packages", app.getClientTravelPackagesHandler)
	router.HandlerFunc(http.MethodPost, "/customers/:id/clients/:clientId/travel-packages", app.createClientTravelPackageHandler)
	router.HandlerFunc(http.MethodGet, "/customers/:id/clients/:clientId/travel-packages/:travelPackageId", app.getClientTravelPackageHandler)
	router.HandlerFunc(http.MethodPost, "/customers/:id/clients/:clientId/deploy", app.deployClientHandler)

	// Public routes
	router.HandlerFunc(http.MethodGet, "/public/client/:clientId", app.getPublicClientHandler)

	// Metrics
	router.Handler(http.MethodGet, "/metrics/", app.authenticateAndAuthorize(expvar.Handler()))

	return app.metrics(app.recoverPanic(router))
}
