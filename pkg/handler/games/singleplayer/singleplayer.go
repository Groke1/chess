package singleplayer

import (
	"common"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"pkg/handler/games"
	"pkg/repository"
	"pkg/service"
)

type SinglePlayer struct {
	repo *repository.Repository
	serv service.Service
	*games.AbstractGame
}

func New(repo *repository.Repository, serv service.Service) *SinglePlayer {
	return &SinglePlayer{
		repo:         repo,
		serv:         serv,
		AbstractGame: games.NewGame(serv),
	}
}

func (s *SinglePlayer) GetGames(w http.ResponseWriter, r *http.Request) {
	games.GetGames(w, r, s.repo.Singleplayer)
}

func (s *SinglePlayer) GetGameById(w http.ResponseWriter, r *http.Request) {
	games.GetGameById(w, r, s.repo.Singleplayer)
}

func (s *SinglePlayer) AddGame(w http.ResponseWriter, r *http.Request) {}

func (s *SinglePlayer) GetByResult(w http.ResponseWriter, r *http.Request) {
	games.GetByResult(w, r, s.repo.Singleplayer)
}

func (s *SinglePlayer) GetByTimeControl(w http.ResponseWriter, r *http.Request) {
	games.GetByTimeControl(w, r, s.repo.Singleplayer)
}

func (s *SinglePlayer) CreateGame(w http.ResponseWriter, r *http.Request) {
	var game common.GameWithEngine
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		logrus.Error(err)
		return
	}
	resp := s.serv.CreateSingleGame(game)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
