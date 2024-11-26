package puzzles

import (
	"github.com/gorilla/mux"
	"net/http"
	jh "pkg/handler/json-helper"
	"pkg/repository"
	"strconv"
)

type Puzzles struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Puzzles {
	return &Puzzles{
		repo: repo,
	}
}

func (p *Puzzles) GetPuzzles(w http.ResponseWriter, r *http.Request) {
	jh.JSONResponse(w, func() (any, error) {
		vars := mux.Vars(r)
		return p.repo.Puzzles.GetPuzzles(r.Context(), vars["username"])
	})
}

func (p *Puzzles) GetPuzzleById(w http.ResponseWriter, r *http.Request) {
	jh.JSONResponse(w, func() (any, error) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			return nil, err
		}
		return p.repo.Puzzles.GetPuzzleById(r.Context(), id)
	})
}

func (p *Puzzles) UpdatePuzzle(w http.ResponseWriter, r *http.Request) {

}

func (p *Puzzles) CreatePuzzle(w http.ResponseWriter, r *http.Request) {

}

func (p *Puzzles) DeletePuzzle(w http.ResponseWriter, r *http.Request) {
}

func (p *Puzzles) SolvedPuzzle(w http.ResponseWriter, r *http.Request) {
	jh.JSONResponse(w, func() (any, error) {
		vars := mux.Vars(r)
		return p.repo.Puzzles.GetSolvedPuzzles(r.Context(), vars["username"])
	})
}

func (p *Puzzles) UnsolvedPuzzle(w http.ResponseWriter, r *http.Request) {
	jh.JSONResponse(w, func() (any, error) {
		vars := mux.Vars(r)
		return p.repo.Puzzles.GetUnsolvedPuzzles(r.Context(), vars["username"])
	})
}

func (p *Puzzles) GetRandomPuzzle(w http.ResponseWriter, r *http.Request) {
	jh.JSONResponse(w, func() (any, error) {
		vars := mux.Vars(r)
		return p.repo.Puzzles.GetRandomPuzzle(r.Context(), vars["username"])
	})
}
