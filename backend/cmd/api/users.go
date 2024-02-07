package main

import (
	"errors"
	"net/http"
)

func (app *application) getUsersHandler(w http.ResponseWriter, r *http.Request) {
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

	db_users, err := app.models.Users.GetAll()

	app.writeJSON(w, http.StatusOK, db_users, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	app.logSuccess(r, http.StatusOK, db_users)
}
