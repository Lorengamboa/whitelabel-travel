package main

import (
	"fmt"
	"errors"
	"net/http"

	"v0/internal/data"
	"v0/internal/validator"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	// Expected data from the user
	var input struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
	}
	// Try reading the user input to JSON
	err := app.readJSON(w, r, &input)
	if err != nil {

		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	// Hash user password
	err = user.Password.Set(input.Password)
	if err != nil {

		app.serverErrorResponse(w, r, err)
		return
	}

	// Validate the user input
	v := validator.New()
	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Save the user in the database
	userId, err := app.models.Users.Insert(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "A user with this email address already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	
	fmt.Print("User created with id ", userId) 
	
	// Respond with success
	app.successResponse(
		w,
		r,
		http.StatusAccepted,
		"Your account creation was accepted successfully.",
	)
}