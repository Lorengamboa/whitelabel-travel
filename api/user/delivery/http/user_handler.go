package http

import (
	"encoding/json"
	"net/http"

	"github.com/Lorengamboa/whitelabel-travel/user/domain"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(mux *http.ServeMux, userUsecase domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: userUsecase,
	}

	mux.Handle("/users/get", http.HandlerFunc(handler.GetAll))
}

func (u *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := u.UserUsecase.GetAll(r.Context())

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
