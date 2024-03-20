package domain

import (
	"context"
	"net/http"

	"github.com/Lorengamboa/whitelabel-travel/user/domain"
)

type Auth struct {
	User   *domain.User `json:"user"`
	Cookie *http.Cookie `json:"_"`
}

type AuthUsecase interface {
	Login(ctx context.Context, email, password string) (*Auth, error)
	Register(ctx context.Context, email, password string) (*Auth, error)
}
