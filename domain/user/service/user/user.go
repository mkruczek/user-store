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

func (s *Service) Create(dto user.PrivateView) (*user.PrivateView, *errors.RestError) {
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

	var result user.PrivateView
	result.FromModel(&e)

	return &result, nil
}

func (s *Service) PartialUpdate(id string, dto user.UpdateDTO) (*user.PrivateView, *errors.RestError) {
	uid, err := parseID(id)
	if err != nil {
		return nil, err
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

	var result user.PrivateView
	result.FromModel(e)

	return &result, nil
}

func (s *Service) FullUpdate(id string, dto user.PrivateView) (*user.PrivateView, *errors.RestError) {
	uid, err := parseID(id)
	if err != nil {
		return nil, err
	}

	oldUser, getErr := s.repo.GetByID(uid)
	if getErr != nil {
		return nil, getErr
	}

	errs := s.validator.CreateUser(&dto) //consider other validation for full update
	if len(errs) > 0 {
		return nil, errors.NewBadRequestErrorValidationList(errs)
	}

	e := user.Model{
		ID:          oldUser.ID,
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		Email:       dto.Email,
		CreatedDate: oldUser.CreatedDate,
		UpdateDate:  time.Now().UTC(),
	}

	errUpdate := s.repo.Update(&e)
	if errUpdate != nil {
		return nil, errUpdate
	}

	var result user.PrivateView
	result.FromModel(&e)

	return &result, nil
}

func (s *Service) GetById(isPublic bool, id string) (user.View, *errors.RestError) {
	uid, err := parseID(id)
	if err != nil {
		return nil, err
	}

	m, getErr := s.repo.GetByID(uid)
	if getErr != nil {
		return nil, getErr
	}

	return m.ToView(isPublic), nil
}

func (s *Service) Search(isPublic bool, values map[string][]string) ([]user.View, *errors.RestError) {
	models, err := s.repo.Search(values)
	if err != nil {
		return nil, err
	}

	result := make([]user.View, len(models))
	for i, m := range models {
		result[i] = m.ToView(isPublic)
	}
	return result, nil
}

func (s *Service) Delete(id string) *errors.RestError {
	uid, err := parseID(id)
	if err != nil {
		return err
	}

	if delErr := s.repo.Delete(uid); delErr != nil {
		return delErr
	}

	return nil
}

func parseID(id string) (*uuid.UUID, *errors.RestError) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.NewBadRequestErrorf("couldn't parse id : %s", id)
	}
	return &uid, nil
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
