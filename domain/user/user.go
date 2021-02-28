package user

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Validateable interface {
}

type Model struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	Email       string
	CreatedDate time.Time
	UpdateDate  time.Time
}
type DTO struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	CreatedDate int64  `json:"createdDate"`
	UpdateDate  int64  `json:"updateDate"`
}

type UpdateDTO struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email"`
}

func (e *Model) ToDTO() *DTO {
	return &DTO{
		ID:          e.ID.String(),
		FirstName:   e.FirstName,
		LastName:    e.LastName,
		FullName:    fmt.Sprintf("%s %s", e.FirstName, e.LastName),
		Email:       e.Email,
		CreatedDate: e.CreatedDate.UnixNano(),
		UpdateDate:  e.UpdateDate.UnixNano(),
	}
}
