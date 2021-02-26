package user

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mkruczek/user-store/datasource/postgresql"
	"github.com/mkruczek/user-store/domain/user"
	"github.com/mkruczek/user-store/utils/errors"
	"strings"
	"time"
)

type DBUserProvider interface {
	Save(u *user.Model) *errors.RestError
	Update(u *user.Model) *errors.RestError
	GetByID(id uuid.UUID) (*user.Model, *errors.RestError)
	Search(values map[string]string) ([]*user.Model, *errors.RestError)
	Delete(id uuid.UUID) *errors.RestError
	CheckEmailExist(email string) (bool, *errors.RestError)
}

type Repository struct {
	db *postgresql.UserDB
}

func NewUserRepository() *Repository {
	return &Repository{
		db: postgresql.NewUserDBConnection(),
	}
}

func (r *Repository) Save(u *user.Model) *errors.RestError {
	stmt, err := r.db.Prepare(`INSERT INTO users (id, first_name, last_name, email, create_date) VALUES ($1, $2, $3, $4, $5);`)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(u.ID, u.FirstName, u.LastName, u.Email, u.CreatedDate.UnixNano())
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	fmt.Println(insertResult)

	return nil
}

func (r *Repository) Update(u *user.Model) *errors.RestError {
	return errors.NewNotImplementingYet("repository.user.Update")
}

func (r *Repository) GetByID(id uuid.UUID) (*user.Model, *errors.RestError) {
	stmt, err := r.db.Prepare(`SELECT id, first_name, last_name, email, create_date FROM users WHERE id=$1`)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	udb := struct {
		id         uuid.UUID
		firstName  string
		lastName   string
		email      string
		createDate int64
	}{}
	row := stmt.QueryRow(id)
	if err := row.Scan(&udb.id, &udb.firstName, &udb.lastName, &udb.email, &udb.createDate); err != nil {
		if strings.EqualFold("sql: no rows in result set", err.Error()) {
			return nil, errors.NewNotFoundErrorf("couldn't find user with id : %s", id.String())
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &user.Model{
		ID:          udb.id,
		FirstName:   udb.firstName,
		LastName:    udb.lastName,
		Email:       udb.email,
		CreatedDate: time.Unix(0, udb.createDate),
	}, nil
}

func (r *Repository) Search(values map[string]string) ([]*user.Model, *errors.RestError) {
	return nil, errors.NewNotImplementingYet("repository.user.Search")
}

func (r *Repository) Delete(id uuid.UUID) *errors.RestError {
	return errors.NewNotImplementingYet("repository.user.Delete")
}

func (r *Repository) CheckEmailExist(email string) (bool, *errors.RestError) {
	stmt, err := r.db.Prepare(`SELECT id FROM users WHERE email=$1`)
	if err != nil {
		return false, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Query(email)
	if err != nil {
		return false, errors.NewInternalServerError(err.Error())
	}

	return result.Next(), nil
}
