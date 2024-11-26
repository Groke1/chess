package pieces

import (
	"pkg/chess-logic/pkg/functions"
)

type Color byte

const (
	White Color = iota
	Black
)

var colorByte = map[Color]byte{
	White: 'w',
	Black: 'b',
}

func GetByteByColor(color Color) byte {
	return colorByte[color]
}

func NewColorByByte(b byte) (Color, error) {
	return functions.FindKeyByValue[Color, byte](colorByte, b)
}

func GetAnotherColor(color Color) Color {
	if color == White {
		return Black
	}
	return White
}
