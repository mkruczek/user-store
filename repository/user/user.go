package user

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mkruczek/user-store/datasource/postgresql"
	"github.com/mkruczek/user-store/domain/user"
	"github.com/mkruczek/user-store/utils/errors"
	"strings"
)

const (
	insertQuery = "INSERT INTO \"user\" (id, first_name, last_name, email, create_date) VALUES ($1, $2, $3, $4, $5);"
)

//currently Repository works based at map
type Repository struct {
	db    *postgresql.UserDB
	dbMap map[uuid.UUID]*user.Entity
}

func NewUserRepository() *Repository {
	return &Repository{
		db:    postgresql.NewUserDBConnection(),
		dbMap: make(map[uuid.UUID]*user.Entity),
	}
}

func (r *Repository) Save(u *user.Entity) *errors.RestError {
	r.dbMap[u.ID] = u

	stmt, err := r.db.Prepare(insertQuery)
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

func (r *Repository) Update(u *user.Entity) *errors.RestError {
	return errors.NewNotImplementingYet("repository.user.Update")
}

func (r *Repository) GetByID(id uuid.UUID) (*user.Entity, *errors.RestError) {
	u, ok := r.dbMap[id]
	if !ok {
		return nil, errors.NewNotFoundError("not found user for id : " + id.String())
	}
	return u, nil
}

func (r *Repository) Search(values map[string]string) ([]*user.Entity, *errors.RestError) {
	return nil, errors.NewNotImplementingYet("repository.user.Search")
}

func (r *Repository) Delete(id uuid.UUID) *errors.RestError {
	return errors.NewNotImplementingYet("repository.user.Delete")
}

func (r *Repository) CheckEmailExist(email string) bool { //todo consider to add to interface
	for _, e := range r.dbMap {
		if strings.EqualFold(e.Email, strings.TrimSpace(email)) {
			return true
		}
	}
	return false
}
