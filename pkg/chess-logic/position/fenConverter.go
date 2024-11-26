package position

import (
	"errors"
	"pkg/chess-logic/pieces"
	"pkg/chess-logic/position/castling"
	"pkg/chess-logic/position/cell"
	"strconv"
	"strings"
)

var incorrectFEN = errors.New("incorrect fen")

func fen2pos(fen string) (*positionImpl, error) {
	elements := strings.Split(fen, " ")
	if len(elements) != 6 {
		return nil, incorrectFEN
	}
	field, err1 := parseField(elements[0])
	if err1 != nil {
		return nil, err1
	}
	whoseMove, err1 := parseWhoseMove(elements[1])
	availCastling, err2 := castling.NewByString(elements[2])
	enPassantCell, err3 := parseEnPassantCell(elements[3], field)
	halfMoveClock, err4 := strconv.Atoi(elements[4])
	fullMoveNumber, err5 := strconv.Atoi(elements[5])
	if err1 != nil || err2 != nil || err3 != nil ||
		err4 != nil || err5 != nil ||
		halfMoveClock < 0 || fullMoveNumber < 0 {
		return nil, incorrectFEN
	}

	p := &positionImpl{
		field:          field,
		whoseMove:      whoseMove,
		castling:       availCastling,
		enPassantCell:  enPassantCell,
		halfMoveClock:  halfMoveClock,
		fullMoveNumber: fullMoveNumber,
		history:        make(map[uint32]int),
	}
	if checkKings(p) != nil {
		return nil, incorrectFEN
	}
	p.history[p.getHash()]++
	p.validMoves = p.getValidMoves(p.whoseMove)
	return p, nil
}

func pos2fen(c *positionImpl) string {
	var fen = strings.Builder{}
	writePos(c, &fen)
	fen.WriteByte(' ')
	fen.WriteByte(pieces.GetByteByColor(c.whoseMove))
	fen.WriteByte(' ')
	fen.WriteString(c.castling.String() + " ")
	fen.WriteString(c.GetNotationByCell(c.enPassantCell) + " ")
	fen.WriteString(strconv.Itoa(c.halfMoveClock) + " ")
	fen.WriteString(strconv.Itoa(c.fullMoveNumber))
	return fen.String()
}

func parseField(stringField string) ([]cell.Cell, error) {
	field := emptyField()
	stringRows := strings.Split(stringField, "/")
	if len(stringRows) != Rows {
		return field, incorrectFEN
	}
	indInField := 0
	for _, stringRow := range stringRows {
		copyIndInField := indInField
		for i := 0; i < len(stringRow); i++ {
			b := stringRow[i]
			if '1' <= b && b <= '8' {
				indInField += int(b) - '0'
				if i+1 < len(stringRow) && '1' <= stringRow[i+1] && stringRow[i+1] <= '8' {
					return nil, incorrectFEN
				}
				continue
			}
			piece, err := pieces.NewPieceByByte(b)
			if err != nil || indInField >= len(field) {
				return nil, incorrectFEN
			}

			field[indInField].SetPiece(piece)

			indInField++
		}
		if indInField-copyIndInField != Columns {
			return field, incorrectFEN
		}
	}
	return field, nil
}

func parseWhoseMove(whoseMove string) (pieces.Color, error) {
	var color pieces.Color
	if len(whoseMove) != 1 {
		return color, incorrectFEN
	}
	return pieces.NewColorByByte(whoseMove[0])
}

func parseEnPassantCell(cell string, field []cell.Cell) (cell.Cell, error) {
	if cell == "-" {
		return nil, nil
	}
	if len(cell) != 2 {
		return nil, incorrectFEN
	}
	row := getRowByByte(cell[1])
	col := getColumnByByte(cell[0])
	if row < 0 || row >= Rows ||
		col < 0 || col >= Columns {
		return nil, incorrectFEN
	}
	return field[row*Rows+col], nil
}

func checkKings(p *positionImpl) error {
	var whiteKing, blackKing cell.Cell
	for row := 0; row < Rows; row++ {
		for col := 0; col < Columns; col++ {
			getCell, _ := p.GetCell(row, col)
			if getCell.Piece() == nil ||
				getCell.Piece().PieceType() != pieces.King {
				continue
			}
			if getCell.Piece().Color() == pieces.White {
				if whiteKing != nil {
					return incorrectFEN
				}
				whiteKing = getCell
				continue
			}
			if blackKing != nil {
				return incorrectFEN
			}
			blackKing = getCell
		}
	}
	if whiteKing == nil || blackKing == nil ||
		p.isBeatCell(pieces.Black, whiteKing) && p.whoseMove == pieces.Black ||
		p.isBeatCell(pieces.White, blackKing) && p.whoseMove == pieces.White {
		return incorrectFEN
	}
	return nil
}

func writePos(c *positionImpl, fen *strings.Builder) {
	for row := 0; row < Rows; row++ {
		skip := 0
		for col := 0; col < Columns; col++ {
			getCell, _ := c.GetCell(row, col)
			if getCell.Piece() == nil {
				skip++
				continue
			}

			if skip != 0 {
				fen.WriteString(strconv.Itoa(skip))
				skip = 0
			}
			byteName := getCell.Piece().GetByteByPiece()
			fen.WriteByte(byteName)
		}
		if skip != 0 {
			fen.WriteString(strconv.Itoa(skip))
		}
		if row != Rows-1 {
			fen.WriteByte('/')
		}
	}
}
