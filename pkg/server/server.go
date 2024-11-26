package server

import (
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func New(port string, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   5 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}
