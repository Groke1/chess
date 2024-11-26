package castling

import (
	"errors"
	"pkg/chess-logic/pieces"
)

var invalidCastlingString = errors.New("invalid castling string")

type Castling interface {
	GetCastlingByColor(color pieces.Color) (bool, bool)
	SetUnavailable(color pieces.Color)
	SetUnavailableSide(color pieces.Color, isKingSide bool)
	String() string
}

type castlingImpl struct {
	whiteKingSideCastling  bool
	whiteQueenSideCastling bool
	blackKingSideCastling  bool
	blackQueenSideCastling bool
}

func (c *castlingImpl) GetCastlingByColor(color pieces.Color) (bool, bool) {
	if color == pieces.White {
		return c.whiteKingSideCastling, c.whiteQueenSideCastling
	}
	return c.blackKingSideCastling, c.blackQueenSideCastling
}

func (c *castlingImpl) SetUnavailable(color pieces.Color) {
	c.SetUnavailableSide(color, false)
	c.SetUnavailableSide(color, true)
}

func (c *castlingImpl) SetUnavailableSide(color pieces.Color, isKingSide bool) {
	if isKingSide {
		if color == pieces.White {
			c.whiteKingSideCastling = false
		} else {
			c.blackKingSideCastling = false
		}
		return
	}
	if color == pieces.White {
		c.whiteQueenSideCastling = false
	} else {
		c.blackQueenSideCastling = false
	}

}

func NewByString(s string) (*castlingImpl, error) {
	if len(s) == 0 || len(s) > 4 {
		return nil, invalidCastlingString
	}
	list := []bool{false, false, false, false}
	bytes := []byte{'K', 'Q', 'k', 'q'}

	if s == "-" {
		return getCastlingByList(list), nil
	}

	ind := 0
	for i := 0; i < len(bytes); i++ {
		if s[ind] != bytes[i] {
			continue
		}
		list[i] = true
		ind++
		if ind >= len(s) {
			break
		}
	}
	if ind != len(s) {
		return nil, invalidCastlingString
	}
	return getCastlingByList(list), nil
}

func getCastlingByList(list []bool) *castlingImpl {
	return &castlingImpl{
		whiteKingSideCastling:  list[0],
		whiteQueenSideCastling: list[1],
		blackKingSideCastling:  list[2],
		blackQueenSideCastling: list[3],
	}
}

func (c *castlingImpl) String() string {
	res := ""
	if c.whiteKingSideCastling {
		res += "K"
	}
	if c.whiteQueenSideCastling {
		res += "Q"
	}
	if c.blackKingSideCastling {
		res += "k"
	}
	if c.blackQueenSideCastling {
		res += "q"
	}
	if len(res) == 0 {
		return "-"
	}

	return res
}
