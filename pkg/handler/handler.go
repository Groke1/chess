package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"pkg/handler/auth"
	"pkg/handler/games"
	"pkg/handler/games/multiplayer"
	"pkg/handler/games/singleplayer"
	"pkg/handler/puzzles"
	"pkg/handler/roles"
	"pkg/handler/stat"
	"pkg/repository"
	"pkg/service"
)

type Handler struct {
	*puzzles.Puzzles
	*multiplayer.Multiplayer
	*singleplayer.SinglePlayer
	*stat.Stat
	*roles.Roles
	*auth.Auth
}

func New(repo *repository.Repository, serv service.Service) *Handler {
	return &Handler{
		Puzzles:      puzzles.New(repo),
		Multiplayer:  multiplayer.New(repo, serv),
		SinglePlayer: singleplayer.New(repo, serv),
		Stat:         stat.New(repo),
		Roles:        roles.New(repo),
		Auth:         auth.New(repo),
	}
}

func gameRoutes(router *mux.Router, gameRoute string, gameType games.Games) {
	gameRouter := router.PathPrefix(gameRoute).Subrouter()
	gameRouter.HandleFunc("", gameType.CreateGame).Methods(http.MethodPost)
	gameUserRouter := gameRouter.PathPrefix("/{username}").Subrouter()
	gameUserRouter.HandleFunc("", gameType.GetGames).Methods(http.MethodGet)
	gameUserRouter.HandleFunc("/{id}", gameType.GetGameById).Methods(http.MethodGet)
	gameUserRouter.HandleFunc("/{result}", gameType.GetByResult).Methods(http.MethodPost)
	gameUserRouter.HandleFunc("/{time_control}", gameType.GetByTimeControl).Methods(http.MethodPost)
}

func (h *Handler) Routes(router *mux.Router) {

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("", h.Auth.AddUser).Methods(http.MethodPost)

	apiRouter := router.PathPrefix("/api").Subrouter()

	puzzleRouter := apiRouter.PathPrefix("/puzzles").Subrouter()
	puzzlesUserRouter := puzzleRouter.PathPrefix("/{username}").Subrouter()
	puzzlesUserRouter.HandleFunc("", h.Puzzles.GetPuzzles).Methods(http.MethodGet)
	puzzleRouter.HandleFunc("/{id}", h.Puzzles.GetPuzzleById).Methods(http.MethodGet)
	puzzleRouter.HandleFunc("", h.Puzzles.CreatePuzzle).Methods(http.MethodPost)
	puzzleRouter.HandleFunc("/{id}", h.Puzzles.UpdatePuzzle).Methods(http.MethodPut)
	puzzleRouter.HandleFunc("/{id}", h.Puzzles.DeletePuzzle).Methods(http.MethodDelete)
	puzzlesUserRouter.HandleFunc("/solved", h.Puzzles.SolvedPuzzle).Methods(http.MethodGet)
	puzzlesUserRouter.HandleFunc("/unsolved", h.Puzzles.UnsolvedPuzzle).Methods(http.MethodGet)

	gameRoutes(apiRouter, "/multi", h.Multiplayer)

	gameRoutes(apiRouter, "/single", h.SinglePlayer)
	apiRouter.HandleFunc("/move", h.SinglePlayer.MakeMove).Methods(http.MethodPost)

	statRouter := apiRouter.PathPrefix("/stat").Subrouter()
	statUserRouter := statRouter.PathPrefix("/{username}").Subrouter()
	statUserRouter.HandleFunc("", h.Stat.GetStat).Methods(http.MethodGet)

	statRouter.HandleFunc("/top/{amount}", h.Stat.GetTop).Methods(http.MethodGet)

	roleRouter := apiRouter.PathPrefix("/roles").Subrouter()
	roleRouter.HandleFunc("", h.Roles.UpdateRole).Methods(http.MethodPut)

}
