package position

import (
	"pkg/chess-logic/move"
	"pkg/chess-logic/pieces"
	"pkg/chess-logic/position/cell"
)

func (p *positionImpl) isValidMove(move move.Move) bool {
	if move.From().Piece() == nil || move.From().Piece().Color() != p.whoseMove {
		return false
	}

	for _, validMove := range p.validMoves {
		if validMove.From() == move.From() && validMove.To() == move.To() {
			return true
		}
	}
	return false

}

func (p *positionImpl) getValidMoves(color pieces.Color) []move.Move {
	moves := make([]move.Move, 0)
	for row := 0; row < Rows; row++ {
		for col := 0; col < Columns; col++ {
			getCell, _ := p.GetCell(row, col)
			if getCell.Piece() == nil || getCell.Piece().Color() != color {
				continue
			}
			p.appendAllMoves(getCell, &moves)
		}
	}
	moves = p.getAvailableMoves(moves, color)
	p.appendCastlingMoves(color, &moves)
	return moves
}

func (p *positionImpl) swapFromTo(from, to cell.Cell, kingCell *cell.Cell) pieces.Piece {
	if from.Piece().PieceType() == pieces.King {
		*kingCell = to
	}
	savePiece := to.Piece()
	to.SetPiece(from.Piece())
	from.SetPiece(nil)
	return savePiece
}

func (p *positionImpl) getAvailableMoves(moves []move.Move, color pieces.Color) []move.Move {
	kingCell := p.findKing(color)
	availableMoves := make([]move.Move, 0, len(moves))
	for _, mv := range moves {
		to, from := mv.To(), mv.From()
		savePiece := p.swapFromTo(from, to, &kingCell)

		if !p.isBeatCell(pieces.GetAnotherColor(color), kingCell) {
			availableMoves = append(availableMoves, mv)
		}

		p.swapFromTo(to, from, &kingCell)
		to.SetPiece(savePiece)
	}
	return availableMoves
}

func (p *positionImpl) findKing(color pieces.Color) cell.Cell {
	for row := 0; row < Rows; row++ {
		for col := 0; col < Columns; col++ {
			getCell, _ := p.GetCell(row, col)
			if getCell.Piece() != nil && getCell.Piece().Color() == color &&
				getCell.Piece().PieceType() == pieces.King {
				return getCell
			}
		}
	}
	panic("there is no king on the board")
}

func (p *positionImpl) appendAllMoves(cell cell.Cell, moves *[]move.Move) {
	switch cell.Piece().PieceType() {
	case pieces.King:
		p.appendKingMoves(cell, moves)
	case pieces.Queen:
		p.appendQueenMoves(cell, moves)
	case pieces.Bishop:
		p.appendBishopMoves(cell, moves)
	case pieces.Rook:
		p.appendRookMoves(cell, moves)
	case pieces.Knight:
		p.appendKnightMoves(cell, moves)
	case pieces.Pawn:
		p.appendPawnMoves(cell, moves)
	default:
		panic("unknown piece type")
	}
}

func (p *positionImpl) appendKingMoves(cell cell.Cell, moves *[]move.Move) {
	p.appendKingKnightMoves(cell, []int{-1, 0, 1}, func(dx, dy int) bool {
		return dx != 0 || dy != 0
	}, moves)
}

func (p *positionImpl) isEmptyAndSafeCells(kingCell cell.Cell, dx int) bool {
	anotherColor := pieces.GetAnotherColor(kingCell.Piece().Color())
	col := kingCell.Col()
	for {
		col += dx
		getCell, _ := p.GetCell(kingCell.Row(), col)
		if getCell.Piece() == nil {
			continue
		}
		if !(getCell.Piece().PieceType() == pieces.Rook &&
			getCell.Piece().Color() == kingCell.Piece().Color()) {
			return false
		}
		break
	}

	col = kingCell.Col()
	for i := 0; i < 3; i++ {
		getCell, _ := p.GetCell(kingCell.Row(), col)
		if p.isBeatCell(anotherColor, getCell) {
			return false
		}
		col += dx
	}
	return true
}

func (p *positionImpl) appendCastlingMoves(color pieces.Color, moves *[]move.Move) {

	kingSide, queenSide := p.castling.GetCastlingByColor(color)
	kingRow := 0
	if color == pieces.White {
		kingRow = Rows - 1
	}
	kingCell, _ := p.GetCell(kingRow, Columns/2)
	if kingSide &&
		p.isEmptyAndSafeCells(kingCell, 1) {
		kingTo, _ := p.GetCell(kingRow, kingCell.Col()+2)
		*moves = append(*moves, move.New(kingCell, kingTo))
	}
	if queenSide &&
		p.isEmptyAndSafeCells(kingCell, -1) {
		kingTo, _ := p.GetCell(kingRow, kingCell.Col()-2)
		*moves = append(*moves, move.New(kingCell, kingTo))
	}
}

func (p *positionImpl) appendQueenMoves(cell cell.Cell, moves *[]move.Move) {
	p.appendBishopMoves(cell, moves)
	p.appendRookMoves(cell, moves)
}

func (p *positionImpl) appendBishopMoves(cell cell.Cell, moves *[]move.Move) {
	p.appendReachableCells(cell, 1, 1, moves)
	p.appendReachableCells(cell, 1, -1, moves)
	p.appendReachableCells(cell, -1, 1, moves)
	p.appendReachableCells(cell, -1, -1, moves)
}

func (p *positionImpl) appendRookMoves(cell cell.Cell, moves *[]move.Move) {
	p.appendReachableCells(cell, 1, 0, moves)
	p.appendReachableCells(cell, 0, 1, moves)
	p.appendReachableCells(cell, -1, 0, moves)
	p.appendReachableCells(cell, 0, -1, moves)
}

func (p *positionImpl) appendKnightMoves(cell cell.Cell, moves *[]move.Move) {
	p.appendKingKnightMoves(cell, []int{-1, 1, 2, -2}, func(dx, dy int) bool {
		return dx != dy && dx != -dy
	}, moves)
}

func (p *positionImpl) appendKingKnightMoves(cell cell.Cell, delta []int, filter func(dx, dy int) bool, moves *[]move.Move) {
	row, col := cell.Row(), cell.Col()
	for _, dx := range delta {
		for _, dy := range delta {
			if !filter(dx, dy) {
				continue
			}
			moveCell, err := p.GetCell(row+dx, col+dy)
			if err != nil ||
				(moveCell.Piece() != nil && moveCell.Piece().Color() == cell.Piece().Color()) {
				continue
			}
			*moves = append(*moves, move.New(cell, moveCell))
		}
	}
}

func (p *positionImpl) isStartPawnPosition(cell cell.Cell) bool {
	if cell.Piece().Color() == pieces.White {
		return cell.Row() == Rows-2
	}
	return cell.Row() == 1
}

func (p *positionImpl) appendPawnMoves(cell cell.Cell, moves *[]move.Move) {
	row, col := cell.Row(), cell.Col()
	dy := -1
	if cell.Piece().Color() == pieces.Black {
		dy = 1
	}
	moveCell, err := p.GetCell(row+dy, col)
	if err == nil && moveCell.Piece() == nil {
		*moves = append(*moves, move.New(cell, moveCell))
		if p.isStartPawnPosition(cell) {
			moveCell, _ = p.GetCell(row+2*dy, col)
			if moveCell.Piece() == nil {
				*moves = append(*moves, move.New(cell, moveCell))
			}
		}
	}
	p.appendPawnEatMoves(cell, moves)
}

func (p *positionImpl) appendPawnEatMoves(cell cell.Cell, moves *[]move.Move) {
	dy := -1
	if cell.Piece().Color() == pieces.Black {
		dy = 1
	}

	p.appendEatPawnMove(cell, 1, dy, moves)
	p.appendEatPawnMove(cell, -1, dy, moves)
}

func (p *positionImpl) appendEatPawnMove(cell cell.Cell, dx, dy int, moves *[]move.Move) {
	eatCell, err := p.GetCell(cell.Row()+dy, cell.Col()+dx)
	if err == nil && (eatCell.Piece() != nil && eatCell.Piece().Color() != cell.Piece().Color() ||
		eatCell == p.enPassantCell) {
		*moves = append(*moves, move.New(cell, eatCell))
	}
}

func (p *positionImpl) appendReachableCells(from cell.Cell, dx, dy int, moves *[]move.Move) {
	row, col := from.Row(), from.Col()
	for {
		row += dx
		col += dy
		getCell, err := p.GetCell(row, col)
		if err != nil {
			return
		}
		if getCell.Piece() == nil {
			*moves = append(*moves, move.New(from, getCell))
			continue
		}

		if getCell.Piece().Color() != from.Piece().Color() {
			*moves = append(*moves, move.New(from, getCell))
		}
		return
	}
}

func (p *positionImpl) isBeatCell(color pieces.Color, c cell.Cell) bool {
	typePieces := []pieces.PieceType{
		pieces.Queen, pieces.Knight, pieces.Bishop,
		pieces.Rook, pieces.Pawn, pieces.King,
	}
	anotherColor := pieces.GetAnotherColor(color)
	for _, typePiece := range typePieces {
		moves := make([]move.Move, 0)

		newC := cell.New(pieces.New(typePiece, anotherColor), c.Row(), c.Col())
		if typePiece == pieces.Pawn {
			p.appendPawnEatMoves(newC, &moves)
		} else {
			p.appendAllMoves(newC, &moves)
		}

		for _, m := range moves {
			to := m.To()
			if to.Piece() != nil && to.Piece().Color() == color && to.Piece().PieceType() == typePiece {
				return true
			}
		}
	}
	return false
}
