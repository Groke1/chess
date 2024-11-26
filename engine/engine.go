package engine

import (
	"sync/atomic"
)

type Engine interface {
	Start() error
	GetMove(fen string, depth int, level int) string
	GetEval(fen string) string
	Close() error
	Lock()
	Unlock()
}

type Pool struct {
	engines   []Engine
	indEngine atomic.Int32
}

func New(engines []Engine) *Pool {
	return &Pool{
		engines: engines,
	}
}

func (p *Pool) GetEngine() Engine {
	p.indEngine.Add(1)
	ind := int(p.indEngine.Load()) % len(p.engines)
	return p.engines[ind]
}

func (p *Pool) Start() error {
	for _, engine := range p.engines {
		if err := engine.Start(); err != nil {
			//p.Close()
			return err
		}
	}
	return nil
}

func (p *Pool) Close() error {
	for _, engine := range p.engines {
		if err := engine.Close(); err != nil {
			return err
		}
	}
	return nil
}
