package cell

import (
	"pkg/chess-logic/pieces"
)

type Cell interface {
	Row() int
	Col() int
	Piece() pieces.Piece
	SetPiece(piece pieces.Piece)
}

type cellImpl struct {
	piece pieces.Piece
	row   int
	col   int
}

func New(piece pieces.Piece, row, col int) *cellImpl {
	return &cellImpl{
		piece: piece,
		row:   row,
		col:   col,
	}
}
func (c *cellImpl) Piece() pieces.Piece {
	return c.piece
}

func (c *cellImpl) SetPiece(piece pieces.Piece) {
	c.piece = piece
}

func (c *cellImpl) Row() int {
	return c.row
}

func (c *cellImpl) Col() int {
	return c.col
}
