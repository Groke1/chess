package pieces

type PieceType byte

const (
	King PieceType = iota
	Queen
	Knight
	Bishop
	Rook
	Pawn
)

var pieceByte = map[PieceType]byte{
	King:   'k',
	Queen:  'q',
	Knight: 'n',
	Bishop: 'b',
	Rook:   'r',
	Pawn:   'p',
}
