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

func (u *Validator) CreateUser(dto *user.PrivateView) []errors.RestError {
	var result []errors.RestError
	err := u.validateEmail(dto.Email, &result)
	if err != nil {
		result = append(result, *err)
		return result
	}

	return result
}

func (u *Validator) UpdateUser(dto *user.UpdateDTO) []errors.RestError {
	var result []errors.RestError

	if dto.Email != nil {
		err := u.validateEmail(*dto.Email, &result)
		if err != nil {
			result = append(result, *err)
			return result
		}
	}

	return result
}

func (u *Validator) validateEmail(email string, result *[]errors.RestError) *errors.RestError {

	if isOk := func(s string) bool {
		email = strings.TrimSpace(email)
		return emailRegex.MatchString(email)
	}(email); !isOk {
		*result = append(*result, *errors.NewBadRequestError("invalid email"))
	}
	if exist, err := u.repo.CheckEmailExist(email); err != nil {
		return err
	} else if exist {
		*result = append(*result, *errors.NewBadRequestError(fmt.Sprintf("email [%s] already exist", email)))
	}

	return nil
}
