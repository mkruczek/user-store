package validators

import (
	"github.com/mkruczek/user-store/domain/user"
	userRepository "github.com/mkruczek/user-store/repository/user"
	"github.com/mkruczek/user-store/utils/errors"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Validator struct {
	repo *userRepository.Repository
}

func NewUserValidator(repo *userRepository.Repository) *Validator {
	return &Validator{repo: repo}
}

func (u *Validator) User(model *user.DTO) error {

	if email := validateEmailStruct(model.Email); !email {
		return errors.NewBadRequestError("invalid email")
	}
	if exist := u.repo.CheckEmailExist(model.Email); exist {
		return errors.NewBadRequestError("email exist")
	}

	return nil
}

func validateEmailStruct(email string) bool {
	email = strings.TrimSpace(email)
	return emailRegex.MatchString(email)
}
