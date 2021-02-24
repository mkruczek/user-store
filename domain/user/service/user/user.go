package user

import (
	"github.com/google/uuid"
	"github.com/mkruczek/user-store/domain/user"
	"github.com/mkruczek/user-store/domain/user/validators"
	userRepository "github.com/mkruczek/user-store/repository/user"
	"github.com/mkruczek/user-store/utils/errors"
	"strings"
	"time"
)

type DBUserProvider interface {
	//Ping() *errors.RestError //(??)maybe add
	Save(u *user.Entity) *errors.RestError
	Update(u *user.Entity) *errors.RestError
	GetByID(id uuid.UUID) (*user.Entity, *errors.RestError)
	Search(values map[string]string) ([]*user.Entity, *errors.RestError)
	Delete(id uuid.UUID) *errors.RestError
}

type Service struct {
	repo      DBUserProvider
	validator *validators.Validator
}

func NewUserService() *Service {
	userRepo := userRepository.NewUserRepository()
	return &Service{
		repo:      userRepo,
		validator: validators.NewUserValidator(userRepo),
	}
}

func (s *Service) CreateUser(dto user.DTO) (*user.DTO, error) {

	err := s.validator.User(&dto)
	if err != nil {
		return nil, err
	}

	e := user.Entity{
		ID:          uuid.New(),
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		Email:       strings.TrimSpace(dto.Email),
		CreatedDate: time.Now().UTC(),
	}

	err = s.repo.Save(&e) //(??)why i need additional check with (*errors.RestError)(nil)
	if err != nil && err != (*errors.RestError)(nil) {
		return nil, err
	}

	return e.ToDTO(), nil
}

func (s *Service) GetById(id string) (*user.DTO, *errors.RestError) {

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.NewBadRequestErrorf("couldn't parse id : %s", id)
	}

	e, getErr := s.repo.GetByID(uid) //TODO improve error
	if getErr != nil {
		return nil, getErr
	}
	return e.ToDTO(), nil
}
