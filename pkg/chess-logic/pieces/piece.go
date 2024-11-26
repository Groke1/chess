package pieces

import (
	"pkg/chess-logic/pkg/functions"
)

type Piece interface {
	PieceType() PieceType
	Color() Color
	GetByteByPiece() byte
}

type pieceImpl struct {
	pieceType PieceType
	color     Color
}

func New(pieceType PieceType, color Color) *pieceImpl {
	return &pieceImpl{pieceType, color}
}

func NewPieceByByte(b byte) (Piece, error) {
	lowerByte := b
	if 'A' <= lowerByte && lowerByte <= 'Z' {
		lowerByte = b - 'A' + 'a'
	}
	pieceType, err := functions.FindKeyByValue[PieceType, byte](pieceByte, lowerByte)
	if err != nil {
		return nil, err
	}
	color := White
	if lowerByte == b {
		color = Black
	}
	return New(pieceType, color), nil
}

func (p *pieceImpl) PieceType() PieceType {
	return p.pieceType
}

func (p *pieceImpl) Color() Color {
	return p.color
}

func (p *pieceImpl) GetByteByPiece() byte {
	b := pieceByte[p.PieceType()]
	if p.Color() == Black {
		return b
	}
	return b - 'a' + 'A'
}
