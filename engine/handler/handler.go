package handler

import (
	"encoding/json"
	"engine"
	"engine/request"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Handler struct {
	enginePool *engine.Pool
}

func New(pool *engine.Pool) *Handler {
	return &Handler{
		enginePool: pool,
	}
}

func (h *Handler) handlerImpl(r *http.Request) *request.Request {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	req := request.New()

	if err := json.Unmarshal(body, &req); err != nil {
		log.Println(err)
	}

	return req
}

func (h *Handler) HandlerMove(w http.ResponseWriter, r *http.Request) {
	req := h.handlerImpl(r)

	eng := h.enginePool.GetEngine()

	eng.Lock()
	resp := eng.GetMove(req.Fen, req.Depth, req.Level)
	eng.Unlock()

	fmt.Fprint(w, resp)
}

func (h *Handler) HandlerEval(w http.ResponseWriter, r *http.Request) {
	req := h.handlerImpl(r)

	eng := h.enginePool.GetEngine()
	eng.Lock()
	resp := eng.GetEval(req.Fen)
	eng.Unlock()

	fmt.Fprint(w, resp)
}
