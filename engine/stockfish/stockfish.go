package stockfish

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

const (
	startCmd       = "/usr/games/stockfish"
	quitCmd        = "quit"
	positionFenCmd = "position fen"
	goCmd          = "go depth"
	evalCmd        = "eval"
	setLevelCmd    = "setoption name Skill Level value"
)

var (
	startingError = errors.New("error starting Stockfish")
)

type Stockfish struct {
	cmd    *exec.Cmd
	stdin  *bufio.Writer
	stdout *bufio.Scanner
	mx     sync.Mutex
}

func (s *Stockfish) Start() error {
	cmd := exec.Command(startCmd)

	stdin, err1 := cmd.StdinPipe()
	stdout, err2 := cmd.StdoutPipe()
	err3 := cmd.Start()
	if err1 != nil || err2 != nil || err3 != nil {
		return startingError
	}
	s.cmd = cmd
	s.stdin = bufio.NewWriter(stdin)
	s.stdout = bufio.NewScanner(stdout)
	return nil
}

func New() *Stockfish {
	return &Stockfish{}
}

func (s *Stockfish) Lock() {
	s.mx.Lock()
}

func (s *Stockfish) Unlock() {
	s.mx.Unlock()
}

func (s *Stockfish) Close() error {
	s.sendMsg(quitCmd)
	return s.cmd.Wait()
}

func (s *Stockfish) GetEval(fen string) string {
	s.sendMsg(positionFenCmd + " " + fen)
	s.sendMsg(evalCmd)
	resp := s.getResponse("Final evaluation", 2)
	return resp
}

func (s *Stockfish) GetMove(fen string, depth int, level int) string {
	s.sendMsg(setLevelCmd + " " + strconv.Itoa(level))
	s.sendMsg(positionFenCmd + " " + fen)
	s.sendMsg(goCmd + " " + strconv.Itoa(depth))

	resp := s.getResponse("bestmove", 1)

	return resp
}

func (s *Stockfish) sendMsg(msg string) {
	fmt.Fprintln(s.stdin, msg)
	s.stdin.Flush()
}

func (s *Stockfish) getResponse(prefix string, ind int) string {
	for s.stdout.Scan() {
		line := s.stdout.Text()
		if strings.HasPrefix(line, prefix) {
			return strings.Fields(line)[ind]
		}
	}
	return ""
}
