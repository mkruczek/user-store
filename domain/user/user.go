package user

import (
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	Email       string
	CreatedDate time.Time
}
type DTO struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	CreatedDate int64  `json:"createdDate"`
}

func (e *Entity) ToDTO() *DTO {
	return &DTO{
		ID:          e.ID.String(),
		FirstName:   e.FirstName,
		LastName:    e.LastName,
		Email:       e.Email,
		CreatedDate: e.CreatedDate.UnixNano(),
	}
}
