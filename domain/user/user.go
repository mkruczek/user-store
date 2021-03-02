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
type PublicView struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName"`
	CreatedDate int64  `json:"createdDate"`
}

type PrivateView struct {
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

type View interface {
	FromModel(m *Model)
}

func (p *PublicView) FromModel(m *Model) {
	p.ID = m.ID.String()
	p.FirstName = m.FirstName
	p.CreatedDate = m.CreatedDate.UnixNano()
}

func (p *PrivateView) FromModel(m *Model) {
	p.ID = m.ID.String()
	p.FirstName = m.FirstName
	p.LastName = m.LastName
	p.FullName = fmt.Sprintf("%s %s", m.FirstName, m.LastName)
	p.Email = m.Email
	p.CreatedDate = m.CreatedDate.UnixNano()
	p.UpdateDate = m.UpdateDate.UnixNano()
}

func (e *Model) ToView(isPublic bool) View {
	if !isPublic {
		return &PrivateView{
			ID:          e.ID.String(),
			FirstName:   e.FirstName,
			LastName:    e.LastName,
			FullName:    fmt.Sprintf("%s %s", e.FirstName, e.LastName),
			Email:       e.Email,
			CreatedDate: e.CreatedDate.UnixNano(),
			UpdateDate:  e.UpdateDate.UnixNano(),
		}
	}
	return &PublicView{
		ID:          e.ID.String(),
		FirstName:   e.FirstName,
		CreatedDate: e.CreatedDate.UnixNano(),
	}

}
