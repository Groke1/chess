package stat

import (
	"github.com/gorilla/mux"
	"net/http"
	jh "pkg/handler/json-helper"
	"pkg/repository"
	"strconv"
)

type Stat struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Stat {
	return &Stat{
		repo: repo,
	}
}

func (s *Stat) GetStat(w http.ResponseWriter, r *http.Request) {
	jh.JSONResponse(w, func() (any, error) {
		vars := mux.Vars(r)
		return s.repo.Stat.GetUserStat(r.Context(), vars["username"])
	})
}

func (s *Stat) GetTop(w http.ResponseWriter, r *http.Request) {
	jh.JSONResponse(w, func() (any, error) {
		vars := mux.Vars(r)
		amount, err := strconv.Atoi(vars["amount"])
		if err != nil {
			return nil, err
		}
		return s.repo.Stat.GetTopStat(r.Context(), amount)
	})
}

func (s *Stat) GetTimeControlsStat(w http.ResponseWriter, r *http.Request) {
	jh.JSONResponse(w, func() (any, error) {
		vars := mux.Vars(r)
		return s.repo.Stat.GetTimeControlsStat(r.Context(), vars["username"])
	})
}
