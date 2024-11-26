package position

import (
	"github.com/stretchr/testify/require"
	"pkg/chess-logic/move"
	"slices"
	"testing"
)

func getMove(mv string, pos Position) move.Move {
	c1, _ := pos.GetCellByNotation(mv[:2])
	c2, _ := pos.GetCellByNotation(mv[2:])
	return move.New(c1, c2)
}

func getCells() []string {
	s1 := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	s2 := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	res := make([]string, 0, 64)
	for _, letter := range s2 {
		for _, number := range s1 {
			res = append(res, letter+number)
		}
	}
	return res
}

func getOtherMoves(moves []string) []string {
	res := make([]string, 0)
	cells := getCells()
	for _, c1 := range cells {
		for _, c2 := range cells {
			if !slices.Contains(moves, c1+c2) {
				res = append(res, c1)
			}
		}
	}
	return res

}

func TestValidMoves(t *testing.T) {
	testCases := []struct {
		fen          string
		validMoves   []string
		invalidMoves []string
	}{

		{
			fen:        "r1bqkb1r/ppp2ppp/5n2/3P4/1n6/2N5/PPP1QPPP/R1B1KBNR b KQkq - 0 7",
			validMoves: []string{"c8e6", "d8e7", "f8e7", "f6e4", "e8d7"},
		},
		{
			fen:          "r1bqk2r/ppp1bppp/3P1n2/8/1n6/2N5/PPP1QPPP/R1B1KBNR b KQkq - 0 8",
			validMoves:   []string{"e8g8", "c7d6", "c7c5", "c7c6", "b4c2", "c8h3", "d8d6", "a7a5"},
			invalidMoves: []string{"e7d6", "f6h7", "d8d5", "c7b6", "e8c8"},
		},

		{
			fen:        "r1b1k2r/ppp1bppp/3q1n2/1Q6/1n6/2N5/PPP2PPP/R1B1KBNR b KQkq - 1 9",
			validMoves: []string{"c7c6", "b4c6", "d6c6", "d6d7", "c8d7", "e8d8", "e8f8", "f6d7"},
		},

		{
			fen:        "r3k2r/pppbbppp/3q1n2/8/8/2N5/PPn1QPPP/R1B1KBNR w KQkq - 0 11",
			validMoves: []string{"e2c2"},
		},

		{
			fen:          "r3k2r/pp1b1ppp/3Q4/8/8/8/PP1B1q1P/R2K1B1R b kq - 3 18",
			validMoves:   []string{"e8c8", "f2d2", "h8f8", "d7g4"},
			invalidMoves: []string{"e8g8", "e8f7", "f2c2"},
		},

		{
			fen:        "3k4/8/3q3b/8/5P2/3NP3/r1RK4/8 w - - 0 1",
			validMoves: []string{"c2a2", "c2b2", "d2c3", "d2c1", "d2d1", "d2e2", "d2e1", "e3e4", "f4f5"},
		},
		{
			fen:          "r3k2r/pp1pp1pp/1qp2p1n/b7/2B5/4P1B1/PPPP1PPP/R2QK1NR b KQkq - 0 1",
			validMoves:   []string{"e8c8"},
			invalidMoves: []string{"e8g8"},
		},
		{
			fen:        "r3k2r/pp1pp1pp/1qp2p1n/4B3/2B5/1Qb1P3/PP1KNPPP/2R4R w kq - 0 2",
			validMoves: []string{"b2c3", "b3c3", "c1c3", "e5c3", "e2c3", "d2d3", "d2c3", "d2d1", "d2c2"},
		},

		{
			fen:        "5K2/1Bpp4/8/3q3B/8/5knR/8/1r5r b - - 0 1",
			validMoves: []string{"f3f4", "f3e4", "f3e3", "f3f2", "f3g2"},
		},

		{
			fen:          "rnbqkb1r/ppppppp1/2B2n2/4N3/7p/4P3/PPPP1PPP/RNBQK2R b KQkq - 1 5",
			validMoves:   []string{"d7c6", "b7c6", "b7b5", "e7e6", "g7g5"},
			invalidMoves: []string{"d7d6", "d7d5", "c7c5", "e7e5", "f7f5", "e8g8"},
		},

		{
			fen:          "rnbqkbnr/ppppp1p1/7p/4Pp2/8/8/PPPP1PPP/RNBQKBNR w KQkq f6 0 3",
			validMoves:   []string{"e5f6", "e5e6"},
			invalidMoves: []string{"e5d6", "e5e7"},
		},

		{
			fen:          "rnbqkbnr/1ppppppp/8/8/pP2P3/5N2/P1PP1PPP/RNBQKB1R b KQkq b3 0 3",
			validMoves:   []string{"a4b3", "a4a3"},
			invalidMoves: []string{"a8a4"},
		},
	}

	for _, tc := range testCases {
		if tc.invalidMoves == nil {
			tc.invalidMoves = getOtherMoves(tc.validMoves)
		}
	}

	for _, tc := range testCases {
		pos, err := NewFromFEN(tc.fen)
		require.NoError(t, err, tc.fen)
		for _, validMove := range tc.validMoves {
			require.True(t, pos.isValidMove(getMove(validMove, pos)), tc.fen, validMove)
		}
		for _, invalidMove := range tc.invalidMoves {
			require.False(t, pos.isValidMove(getMove(invalidMove, pos)), tc.fen, invalidMove)
		}
	}
}
