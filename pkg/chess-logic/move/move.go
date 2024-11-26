package move

import (
	"pkg/chess-logic/pieces"
	"pkg/chess-logic/position/cell"
)

type Move interface {
	From() cell.Cell
	To() cell.Cell
}

type moveImpl struct {
	from cell.Cell
	to   cell.Cell
}

func New(from, to cell.Cell) *moveImpl {
	return &moveImpl{
		from: from,
		to:   to,
	}
}

func (m *moveImpl) From() cell.Cell {
	return m.from
}

func (m *moveImpl) To() cell.Cell {
	return m.to
}

type PawnTransformMove struct {
	moveImpl
	transformPieceType pieces.PieceType
}

func NewTransformMove(from, to cell.Cell, pieceType pieces.PieceType) *PawnTransformMove {
	return &PawnTransformMove{
		moveImpl:           moveImpl{from, to},
		transformPieceType: pieceType,
	}
}

func (p *PawnTransformMove) TransformPiece() pieces.PieceType {
	return p.transformPieceType
}
