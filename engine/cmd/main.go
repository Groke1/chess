package main

import (
	"engine/handler"
	"engine/stockfish"
	"log"
	"net/http"
)

const (
	Port     = "8080"
	PoolSize = 5
)

func main() {
	pool := stockfish.NewPool(PoolSize)
	if err := pool.Start(); err != nil {
		log.Fatal(err)
	}
	defer func() { _ = pool.Close() }()
	handle := handler.New(pool)
	listenPort(Port, handle)
}

func listenPort(port string, handle *handler.Handler) {
	http.HandleFunc("/move", handle.HandlerMove)
	http.HandleFunc("/eval", handle.HandlerEval)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
