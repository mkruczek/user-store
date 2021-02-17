package user

import (
	"github.com/google/uuid"
	"github.com/mkruczek/user-store/domain/user"
	"github.com/mkruczek/user-store/utils/errors"
	"strings"
)

type Repository struct {
	db map[uuid.UUID]*user.Entity
}

func NewUserRepository() *Repository {
	return &Repository{db: make(map[uuid.UUID]*user.Entity)}
}

func (r *Repository) Save(u *user.Entity) error {
	r.db[u.ID] = u
	return nil
}

func (r *Repository) GetById(id uuid.UUID) (*user.Entity, *errors.RestError) {
	u, ok := r.db[id]
	if !ok {
		return nil, errors.NewNotFoundError("not found user for id : " + id.String())
	}
	return u, nil
}

func (r *Repository) CheckEmailExist(email string) bool {
	for _, e := range r.db {
		if strings.EqualFold(e.Email, strings.TrimSpace(email)) {
			return true
		}
	}
	return false
}
