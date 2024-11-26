package player

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pkg/chess-logic/position"
)

type Player interface {
	Move(p position.Position) string
}

type HumanPlayer struct {
}

func (h *HumanPlayer) Move(p position.Position) string {
	fmt.Println("Move: ")
	var stringMove string
	fmt.Scan(&stringMove)
	return stringMove
}

type StockfishPlayer struct {
	level int
}

func NewStockfishPlayer(level ...int) *StockfishPlayer {
	finalLevel := 20
	if len(level) > 0 {
		finalLevel = level[0]
	}
	return &StockfishPlayer{
		level: finalLevel,
	}

}

type Request struct {
	Fen   string `json:"fen"`
	Depth int    `json:"depth,omitempty"`
	Level int    `json:"level,omitempty"`
}

func (s *StockfishPlayer) Move(p position.Position) string {
	url := "http://stockfish:8080/move"
	r := Request{
		Fen:   p.GetFEN(),
		Depth: 16,
		Level: s.level,
	}
	jsonData, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	stringMove := string(body)
	return stringMove
}
