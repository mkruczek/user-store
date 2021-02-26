package user

import (
	"github.com/mkruczek/user-store/domain/user/validators"
	userRepository "github.com/mkruczek/user-store/repository/user"
)

func NewTestService() *Service {
	userRepo := userRepository.NewMockRepository()
	return &Service{
		repo:      userRepo,
		validator: validators.NewUserValidator(userRepo),
	}
}
