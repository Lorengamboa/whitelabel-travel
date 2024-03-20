package usecase

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Lorengamboa/whitelabel-travel/auth/domain"
	userDomain "github.com/Lorengamboa/whitelabel-travel/user/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userUsecase    userDomain.UserUsecase
	contextTimeout time.Duration
}

func NewAuthUsecase(userUsecase userDomain.UserUsecase, contextTimeout time.Duration) domain.AuthUsecase {
	return &authUsecase{
		userUsecase:    userUsecase,
		contextTimeout: contextTimeout,
	}
}

func (a *authUsecase) createCookie(userId uuid.UUID) (*http.Cookie, error) {
	var buf bytes.Buffer

	err := gob.NewEncoder(&buf).Encode(&userId)
	if err != nil {
		return nil, err
	}

	session := buf.String()

	maxAge, err := strconv.Atoi(os.Getenv("SESSION_EXPIRATION"))
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     "sessionid",
		Value:    session,
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	return cookie, nil
}

func (a *authUsecase) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (a *authUsecase) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a *authUsecase) Login(ctx context.Context, email, password string) (*domain.Auth, error) {
	user, err := a.userUsecase.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if isValid := a.checkPasswordHash(password, user.Password); !isValid {
		return nil, errors.New("wrong credentials")
	}

	cookie, err := a.createCookie(user.ID)
	if err != nil {
		return nil, err
	}

	return &domain.Auth{
		User:   user,
		Cookie: cookie,
	}, nil
}
func (a *authUsecase) Register(ctx context.Context, email, password string) (*domain.Auth, error) { // TODO: handle first name and last name
	user, _ := a.userUsecase.GetByEmail(ctx, email)

	if user != nil {
		return nil, errors.New("user already exists")
	}

	hash, err := a.hashPassword(password)
	if err != nil {
		return nil, err
	}

	user = &userDomain.User{
		Email:    email,
		Password: hash,
	}

	err = a.userUsecase.Insert(ctx, user)
	if err != nil {
		return nil, err
	}

	return &domain.Auth{
		User: user,
	}, nil
}
