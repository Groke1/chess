package games

import (
	"common"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	jh "pkg/handler/json-helper"
	pf "pkg/handler/parse-functions"
	"pkg/repository/games"
	"pkg/service"
	"strconv"
)

type Games interface {
	GetGames(w http.ResponseWriter, r *http.Request)
	GetGameById(w http.ResponseWriter, r *http.Request)
	AddGame(w http.ResponseWriter, r *http.Request)
	GetByResult(w http.ResponseWriter, r *http.Request)
	GetByTimeControl(w http.ResponseWriter, r *http.Request)
	CreateGame(w http.ResponseWriter, r *http.Request)
	MakeMove(w http.ResponseWriter, r *http.Request)
}

type AbstractGame struct {
	serv service.Service
}

func NewGame(serv service.Service) *AbstractGame {
	return &AbstractGame{serv: serv}
}

func (a *AbstractGame) MakeMove(w http.ResponseWriter, r *http.Request) {
	var move common.Move
	if err := json.NewDecoder(r.Body).Decode(&move); err != nil {
		logrus.Error(err)
		return
	}
	resp := a.serv.MakeMove(r.Context(), move)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func GetGames(w http.ResponseWriter, r *http.Request, games games.TypeGames) {
	jh.JSONResponse(w, func() (any, error) {
		vars := mux.Vars(r)
		return games.GetGames(r.Context(), vars["username"])
	})
}

func GetGameById(w http.ResponseWriter, r *http.Request, games games.TypeGames) {
	jh.JSONResponse(w, func() (any, error) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			return nil, err
		}
		return games.GetGameById(r.Context(), id)
	})
}

func GetByResult(w http.ResponseWriter, r *http.Request, games games.TypeGames) {
	jh.JSONResponse(w, func() (any, error) {
		vars := mux.Vars(r)
		return games.GetGamesByResult(r.Context(), vars["username"], vars["result"])
	})
}

func GetByTimeControl(w http.ResponseWriter, r *http.Request, games games.TypeGames) {
	jh.JSONResponse(w, func() (any, error) {
		vars := mux.Vars(r)
		minutes, increment, err := pf.ParseTimeControl(vars["time_control"])
		if err != nil {
			return nil, err
		}
		return games.GetByTimeControl(r.Context(), vars["username"], minutes, increment)
	})
}
