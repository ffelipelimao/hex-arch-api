package usecases

import (
	"context"

	"github.com/ffelipelimao/hex-arch-api/internal/core/domain"
	"github.com/ffelipelimao/hex-arch-api/internal/core/ports"
)

type UserCreator struct {
	userRepository ports.UserRepository
}

func NewUserCreator(userRepository ports.UserRepository) *UserCreator {
	return &UserCreator{
		userRepository: userRepository,
	}
}

func (uc *UserCreator) Create(ctx context.Context, user *domain.User) error {

	err := user.Validate()
	if err != nil {
		return err
	}

	return uc.userRepository.Save(ctx, user)
}
