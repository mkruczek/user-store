package validators

import (
	"fmt"
	"github.com/mkruczek/user-store/domain/user"
	userRepository "github.com/mkruczek/user-store/repository/user"
	"github.com/mkruczek/user-store/utils/errors"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Validator struct {
	repo userRepository.DBUserProvider
}

func NewUserValidator(repo userRepository.DBUserProvider) *Validator {
	return &Validator{repo: repo}
}

func (u *Validator) User(model *user.DTO) *errors.RestError {

	if email := validateEmailStruct(model.Email); !email {
		return errors.NewBadRequestError("invalid email")
	}
	if exist, err := u.repo.CheckEmailExist(model.Email); err != nil {
		return err
	} else if exist {
		return errors.NewBadRequestError(fmt.Sprintf("email [%s] already exist", model.Email))
	}

	return nil
}

func validateEmailStruct(email string) bool {
	email = strings.TrimSpace(email)
	return emailRegex.MatchString(email)
}
