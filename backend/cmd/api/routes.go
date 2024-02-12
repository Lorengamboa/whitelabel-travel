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
	router.HandlerFunc(http.MethodGet, "/customers", app.getCustomersHandler)
	router.HandlerFunc(http.MethodPost, "/customers", app.createCustomerHandler)
	router.HandlerFunc(http.MethodGet, "/customers/:id", app.getCustomerHandler)

	// Metrics
	router.Handler(http.MethodGet, "/metrics/", app.authenticateAndAuthorize(expvar.Handler()))

	return app.metrics(app.recoverPanic(router))
}
