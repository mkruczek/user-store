package user

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mkruczek/user-store/domain/user"
	"time"
)

type Service struct {
}

func NewUserService() *Service {
	return &Service{}
}

func (s *Service) CreateUser(dto user.DTO) (*user.DTO, error) {

	u := user.Entity{
		ID:          uuid.New(),
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		Email:       dto.Email,
		CreatedDate: time.Now(),
	}

	fmt.Printf("save user entity : %v\n", u)

	dto.ID = u.ID.String()
	dto.CreatedDate = u.CreatedDate.Unix()

	return &dto, nil
}
