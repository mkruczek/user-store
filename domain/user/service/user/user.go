package user

import (
	"github.com/google/uuid"
	"github.com/mkruczek/user-store/config"
	"github.com/mkruczek/user-store/domain/user"
	"github.com/mkruczek/user-store/domain/user/validators"
	userRepository "github.com/mkruczek/user-store/repository/user"
	"github.com/mkruczek/user-store/utils/errors"
	"strings"
	"time"
)

type Service struct {
	repo      userRepository.DBUserProvider
	validator *validators.Validator
}

func NewUserService(cfg *config.Config) *Service {
	userRepo := userRepository.NewUserRepository(cfg)
	return &Service{
		repo:      userRepo,
		validator: validators.NewUserValidator(userRepo),
	}
}

func (s *Service) CreateUser(dto user.DTO) (*user.DTO, *errors.RestError) {

	errs := s.validator.CreateUser(&dto)
	if len(errs) > 0 {
		return nil, errors.NewBadRequestErrorValidationList(errs)
	}

	e := user.Model{
		ID:          uuid.New(),
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		Email:       strings.TrimSpace(dto.Email),
		CreatedDate: time.Now().UTC(),
		UpdateDate:  time.Now().UTC(),
	}

	errSave := s.repo.Save(&e)
	if errSave != nil {
		return nil, errSave
	}

	return e.ToDTO(), nil
}

func (s *Service) UpdateUser(id string, dto user.UpdateDTO) (*user.DTO, *errors.RestError) {

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.NewBadRequestErrorf("couldn't parse id : %s", id)
	}

	oldUser, getErr := s.repo.GetByID(uid)
	if getErr != nil {
		return nil, getErr
	}

	errs := s.validator.UpdateUser(&dto)
	if len(errs) > 0 {
		return nil, errors.NewBadRequestErrorValidationList(errs)
	}

	e := updateUserBody(oldUser, dto)

	errUpdate := s.repo.Update(e)
	if errUpdate != nil {
		return nil, errUpdate
	}

	return e.ToDTO(), nil
}

func (s *Service) GetById(id string) (*user.DTO, *errors.RestError) {

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.NewBadRequestErrorf("couldn't parse id : %s", id)
	}

	e, getErr := s.repo.GetByID(uid)
	if getErr != nil {
		return nil, getErr
	}
	return e.ToDTO(), nil
}

func (s *Service) Search(values map[string][]string) ([]*user.DTO, *errors.RestError) {
	models, err := s.repo.Search(values)
	if err != nil {
		return nil, err
	}

	result := make([]*user.DTO, len(models))
	for i, m := range models {
		result[i] = m.ToDTO()
	}
	return result, nil
}

func updateUserBody(old *user.Model, dto user.UpdateDTO) *user.Model {
	e := user.Model{
		ID:          old.ID,
		CreatedDate: old.CreatedDate,
		UpdateDate:  time.Now().UTC(),
	}

	if dto.FirstName != nil {
		e.FirstName = *dto.FirstName
	}

	if dto.LastName != nil {
		e.LastName = *dto.LastName
	}

	if dto.Email != nil {
		e.Email = *dto.Email
	}

	return &e
}
