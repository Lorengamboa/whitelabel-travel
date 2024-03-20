package http

import (
	"encoding/json"
	"net/http"

	"github.com/Lorengamboa/whitelabel-travel/auth/domain"
	userDomain "github.com/Lorengamboa/whitelabel-travel/user/domain"
)

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
}

func NewAuthHandler(mux *http.ServeMux, authUsecase domain.AuthUsecase) {
	handler := &AuthHandler{
		AuthUsecase: authUsecase,
	}

	mux.Handle("/users/login", http.HandlerFunc(handler.Login))
	mux.Handle("/users/register", http.HandlerFunc(handler.Register))
}

func (u *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user *userDomain.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	auth, err := u.AuthUsecase.Login(r.Context(), user.Email, user.Password)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// TODO: to implement
	// err = cookies.WriteEncrypted(w, cookie, app.config.secret.secretKey)
	// if err != nil {
	// 	app.serverErrorResponse(w, r, errors.New("something happened setting your cookie data"))
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auth)
}

func (u *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user *userDomain.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	auth, err := u.AuthUsecase.Register(r.Context(), user.Email, user.Password)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auth)
}
