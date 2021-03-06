package user

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/mkruczek/user-store/config"
	"github.com/mkruczek/user-store/datasource/postgresql"
	"github.com/mkruczek/user-store/domain/user"
	"github.com/mkruczek/user-store/utils/errors"
	"strings"
	"time"
)

type DBUserProvider interface {
	Save(u *user.Model) *errors.RestError
	Update(u *user.Model) *errors.RestError
	GetByID(id *uuid.UUID) (*user.Model, *errors.RestError)
	Search(values map[string][]string) ([]*user.Model, *errors.RestError)
	Delete(id *uuid.UUID) *errors.RestError
	CheckEmailExist(email string) (bool, *errors.RestError)
}

type dbUser struct {
	id         uuid.UUID
	firstName  string
	lastName   string
	email      string
	createDate int64
	updateDate int64
}

type Repository struct {
	db *postgresql.UserDB
}

func NewUserRepository(cfg *config.Config) *Repository {
	return &Repository{
		db: postgresql.NewUserDBConnection(cfg),
	}
}

func (r *Repository) Save(u *user.Model) *errors.RestError {
	stmt, err := r.db.Prepare(`INSERT INTO users (id, first_name, last_name, email, create_date, update_date) VALUES ($1, $2, $3, $4, $5, $6);`)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ID, u.FirstName, u.LastName, u.Email, u.CreatedDate.UnixNano(), u.UpdateDate.UnixNano())
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *Repository) Update(u *user.Model) *errors.RestError {
	stmt, err := r.db.Prepare(`UPDATE users SET first_name = $1, last_name=$2, email=$3, update_date = $4 WHERE id = $5;`)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.UpdateDate.UnixNano(), u.ID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *Repository) GetByID(id *uuid.UUID) (*user.Model, *errors.RestError) {
	stmt, err := r.db.Prepare(`SELECT id, first_name, last_name, email, create_date, update_date FROM users WHERE id=$1`)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	dbu := dbUser{}
	row := stmt.QueryRow(id)
	if err := row.Scan(&dbu.id, &dbu.firstName, &dbu.lastName, &dbu.email, &dbu.createDate, &dbu.updateDate); err != nil {
		if strings.EqualFold("sql: no rows in result set", err.Error()) {
			return nil, errors.NewNotFoundErrorf("couldn't find user with id : %s", id.String())
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return dbu.toModel(), nil
}

func (r *Repository) Search(values map[string][]string) ([]*user.Model, *errors.RestError) {
	sqlQuery := prepareQuery(values)
	stmt, err := r.db.Prepare(sqlQuery)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	var result []*user.Model
	var rows *sql.Rows
	if rows, err = stmt.Query(); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		dbu := dbUser{}
		if err := rows.Scan(&dbu.id, &dbu.firstName, &dbu.lastName, &dbu.email, &dbu.createDate, &dbu.updateDate); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		result = append(result, dbu.toModel())
	}

	return result, nil
}

func (r *Repository) Delete(id *uuid.UUID) *errors.RestError {
	stmt, err := r.db.Prepare(`DELETE FROM users WHERE id=$1`)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *Repository) CheckEmailExist(email string) (bool, *errors.RestError) {
	stmt, err := r.db.Prepare(`SELECT id FROM users WHERE email=$1`)
	if err != nil {
		return false, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	var id uuid.UUID
	row := stmt.QueryRow(email)
	if err := row.Scan(&id); err != nil {
		if strings.EqualFold("sql: no rows in result set", err.Error()) {
			return false, nil
		}
		return false, errors.NewInternalServerError(err.Error())
	}

	return true, nil
}

func prepareQuery(values map[string][]string) string {
	sqlQuery := `SELECT id,first_name, last_name, email, create_date, update_date FROM users WHERE `
	for k, s := range values {
		for i, v := range s {
			sqlQuery += k + "="
			sqlQuery += "'" + v + "'"
			if i < len(s)-1 {
				sqlQuery += " OR "
			}
		}
		sqlQuery += " AND "
	}
	sqlQuery = sqlQuery[:len(sqlQuery)-5]

	return sqlQuery
}

func (d *dbUser) toModel() *user.Model {
	return &user.Model{
		ID:          d.id,
		FirstName:   d.firstName,
		LastName:    d.lastName,
		Email:       d.email,
		CreatedDate: time.Unix(0, d.createDate),
		UpdateDate:  time.Unix(0, d.updateDate),
	}
}
