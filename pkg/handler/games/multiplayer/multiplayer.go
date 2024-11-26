package multiplayer

import (
	"net/http"
	"pkg/handler/games"
	"pkg/repository"
	"pkg/service"
)

type Multiplayer struct {
	repo *repository.Repository
	serv service.Service
	*games.AbstractGame
}

func New(repo *repository.Repository, serv service.Service) *Multiplayer {
	return &Multiplayer{
		repo:         repo,
		serv:         serv,
		AbstractGame: games.NewGame(serv),
	}
}

func (m *Multiplayer) GetGames(w http.ResponseWriter, r *http.Request) {
	games.GetGames(w, r, m.repo.Multiplayer)
}

func (m *Multiplayer) GetGameById(w http.ResponseWriter, r *http.Request) {
	games.GetGameById(w, r, m.repo.Multiplayer)
}

func (m *Multiplayer) AddGame(w http.ResponseWriter, r *http.Request) {}

func (m *Multiplayer) GetByResult(w http.ResponseWriter, r *http.Request) {
	games.GetByResult(w, r, m.repo.Multiplayer)
}

func (m *Multiplayer) GetByTimeControl(w http.ResponseWriter, r *http.Request) {
	games.GetByTimeControl(w, r, m.repo.Multiplayer)
}

func (m *Multiplayer) CreateGame(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
