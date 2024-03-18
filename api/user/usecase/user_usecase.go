package usecase

import (
	"context"
	"time"

	"github.com/Lorengamboa/whitelabel-travel/user/domain"
	"github.com/google/uuid"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepo domain.UserRepository, contextTimeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       userRepo,
		contextTimeout: contextTimeout,
	}
}

func (u *userUsecase) GetById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := u.userRepo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (u *userUsecase) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := u.userRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
