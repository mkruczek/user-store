package user

import (
	"github.com/google/uuid"
	"github.com/mkruczek/user-store/domain/user"
	"github.com/mkruczek/user-store/utils/errors"
	"strings"
)

type MockRepository struct {
	dbMap map[uuid.UUID]*user.Model
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		dbMap: make(map[uuid.UUID]*user.Model),
	}
}

func (r *MockRepository) Save(u *user.Model) *errors.RestError {
	r.dbMap[u.ID] = u
	return nil
}

func (r *MockRepository) Update(u *user.Model) *errors.RestError {
	return nil
}

func (r *MockRepository) GetByID(id uuid.UUID) (*user.Model, *errors.RestError) {
	u, ok := r.dbMap[id]
	if !ok {
		return nil, nil
	}
	return u, nil
}

func (r *MockRepository) Search(values map[string]string) ([]*user.Model, *errors.RestError) {
	return nil, nil
}

func (r *MockRepository) Delete(id uuid.UUID) *errors.RestError {
	return nil
}

func (r *MockRepository) CheckEmailExist(email string) (bool, *errors.RestError) {
	for _, e := range r.dbMap {
		if strings.EqualFold(e.Email, strings.TrimSpace(email)) {
			return true, nil
		}
	}
	return false, nil
}
