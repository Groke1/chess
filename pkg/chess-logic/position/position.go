package position

import "C"
import (
	"errors"
	"hash/fnv"
	"pkg/chess-logic/move"
	"pkg/chess-logic/pieces"
	"pkg/chess-logic/pkg/functions"
	"pkg/chess-logic/position/castling"
	"pkg/chess-logic/position/cell"
	"pkg/chess-logic/result"
	"strings"
)

type Position interface {
	MakeMove(move move.Move) (result.Result, error)
	MakeMoveByString(mv string) (result.Result, error)
	GetFEN() string
	GetCell(row, col int) (cell.Cell, error)
	GetCellByNotation(notation string) (cell.Cell, error)
	String() string
}

var invalidMoveError = errors.New("invalid move")
var indexOutOfBoundsError = errors.New("index out of bounds")

const startPosition = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

const (
	Rows    = 8
	Columns = 8
)

type positionImpl struct {
	field          []cell.Cell
	whoseMove      pieces.Color
	castling       castling.Castling
	enPassantCell  cell.Cell
	halfMoveClock  int
	fullMoveNumber int
	validMoves     []move.Move
	history        map[uint32]int
	moves          []move.Move
}

func NewStartPosition() *positionImpl {
	c, _ := NewFromFEN(startPosition)
	return c
}

func emptyField() []cell.Cell {
	field := make([]cell.Cell, Rows*Columns)
	for ind := 0; ind < Rows*Columns; ind++ {
		field[ind] = cell.New(nil, ind/Columns, ind%Columns)
	}
	return field
}

func NewFromFEN(fen string) (*positionImpl, error) {
	return fen2pos(fen)
}

func (p *positionImpl) GetFEN() string {
	return pos2fen(p)
}

func (p *positionImpl) GetCell(row, col int) (cell.Cell, error) {
	if row >= Rows || row < 0 || col >= Columns || col < 0 {
		return nil, indexOutOfBoundsError
	}

	return p.field[row*Columns+col], nil
}

func (p *positionImpl) changeWhoseMove() {
	p.whoseMove = pieces.GetAnotherColor(p.whoseMove)
}

func (p *positionImpl) MakeMove(mv move.Move) (result.Result, error) {
	if !p.isValidMove(mv) {
		return result.Unknown, invalidMoveError
	}
	p.moves = append(p.moves, mv)

	from := mv.From()
	to := mv.To()

	p.updateEnPassant(mv)

	p.updateCastling(mv)

	p.halfMoveClock++
	if from.Piece().PieceType() == pieces.Pawn ||
		to.Piece() != nil {
		p.halfMoveClock = 0
	}

	if p.whoseMove == pieces.Black {
		p.fullMoveNumber++
	}

	p.checkCastling(mv)

	switch moveType := mv.(type) {
	case *move.PawnTransformMove:
		to.SetPiece(pieces.New(moveType.TransformPiece(), from.Piece().Color()))
	default:
		to.SetPiece(from.Piece())
	}
	from.SetPiece(nil)

	p.changeWhoseMove()
	p.validMoves = p.getValidMoves(p.whoseMove)

	if p.isDraw() {
		return result.Draw, nil
	}

	if p.isWin() {
		return result.Win, nil
	}

	return result.Unknown, nil
}

func (p *positionImpl) MakeMoveByString(mv string) (result.Result, error) {
	if len(mv) != 4 && len(mv) != 5 {
		return result.Unknown, invalidMoveError
	}
	c1, err1 := p.GetCellByNotation(mv[:2])
	c2, err2 := p.GetCellByNotation(mv[2:4])

	if err1 != nil || err2 != nil {
		return result.Unknown, invalidMoveError
	}
	if len(mv) == 5 {

		piece, err3 := pieces.NewPieceByByte(mv[4])
		if err3 != nil || (piece.PieceType() == pieces.King || piece.PieceType() == pieces.Pawn) {
			return result.Unknown, invalidMoveError
		}

		return p.MakeMove(move.NewTransformMove(c1, c2, piece.PieceType()))

	}
	return p.MakeMove(move.New(c1, c2))
}

func (p *positionImpl) updateCastling(move move.Move) {
	pieceType := move.From().Piece().PieceType()
	if pieceType != pieces.Rook && pieceType != pieces.King {
		return
	}
	if pieceType == pieces.King {
		p.castling.SetUnavailable(p.whoseMove)
		return
	}
	p.castling.SetUnavailableSide(p.whoseMove, move.From().Col() > Columns/2)

}

func (p *positionImpl) checkCastling(move move.Move) {
	from, to := move.From(), move.To()
	if from.Piece().PieceType() == pieces.King &&
		functions.AbsInt(from.Col()-to.Col()) == 2 {
		var rookFrom, rookTo cell.Cell
		if to.Col()-from.Col() == 2 {
			rookFrom, _ = p.GetCell(from.Row(), Columns-1)
			rookTo, _ = p.GetCell(from.Row(), to.Col()-1)
		} else {
			rookFrom, _ = p.GetCell(from.Row(), 0)
			rookTo, _ = p.GetCell(from.Row(), to.Col()+1)
		}
		rookTo.SetPiece(rookFrom.Piece())
		rookFrom.SetPiece(nil)
	}
}

func (p *positionImpl) updateEnPassant(mv move.Move) {
	from, to := mv.From(), mv.To()
	if to == p.enPassantCell {
		if p.whoseMove == pieces.White {
			c, _ := p.GetCell(to.Row()+1, to.Col())
			c.SetPiece(nil)
		} else {
			c, _ := p.GetCell(to.Row()-1, to.Col())
			c.SetPiece(nil)
		}
	}

	p.enPassantCell = nil
	if from.Piece().PieceType() == pieces.Pawn &&
		functions.AbsInt(from.Row()-to.Row()) == 2 {
		p.enPassantCell, _ = p.GetCell((from.Row()+to.Row())/2, from.Col())
	}
}

func (p *positionImpl) isDraw() bool {
	hash := p.getHash()
	p.history[hash]++
	return p.halfMoveClock == 50 || p.history[hash] == 3 ||
		(len(p.validMoves) == 0 && !p.isBeatCell(pieces.GetAnotherColor(p.whoseMove), p.findKing(p.whoseMove)))
}

func (p *positionImpl) isWin() bool {
	return len(p.validMoves) == 0
}

func (p *positionImpl) String() string {
	var builder strings.Builder
	for row := 0; row < Rows; row++ {
		for column := 0; column < Columns; column++ {
			getCell, _ := p.GetCell(row, column)
			if getCell.Piece() == nil {
				builder.WriteByte('-')
			} else {
				builder.WriteByte(getCell.Piece().GetByteByPiece())
			}
		}
		builder.WriteByte('\n')
	}
	return builder.String()
}

func getRowByByte(b byte) int {
	return Rows - int(b-'0')
}

func getColumnByByte(b byte) int {
	return int(b - 'a')
}

func (p *positionImpl) GetCellByNotation(notation string) (cell.Cell, error) {
	if len(notation) != 2 {
		return nil, errors.New("incorrect notation")
	}
	return p.GetCell(getRowByByte(notation[1]), getColumnByByte(notation[0]))
}

func (p *positionImpl) GetNotationByCell(cell cell.Cell) string {
	if cell == nil {
		return "-"
	}
	res := string(byte(cell.Col()) + 'a')
	res += string(byte(Rows-cell.Row()) + '0')
	return res
}

func (p *positionImpl) getHash() uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(p.String()))
	return hash.Sum32()
}
