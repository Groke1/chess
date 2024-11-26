package roles

import (
	"net/http"
	"pkg/repository"
)

type Roles struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Roles {
	return &Roles{
		repo: repo,
	}
}

func (role *Roles) UpdateRole(w http.ResponseWriter, r *http.Request) {}
