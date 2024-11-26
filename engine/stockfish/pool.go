package stockfish

import (
	"engine"
)

type Pool struct {
	engine.Pool
}

func NewPool(amount int) *engine.Pool {
	engines := make([]engine.Engine, amount)
	for i := 0; i < amount; i++ {
		engines[i] = New()
	}
	return engine.New(engines)
}
