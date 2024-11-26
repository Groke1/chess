package auth

import (
	"common"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"pkg/repository"
)

type Auth struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Auth {
	return &Auth{
		repo: repo,
	}
}

func (a *Auth) AddUser(w http.ResponseWriter, r *http.Request) {
	var user common.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	a.repo.AddUser(r.Context(), user.Username)

}
