package position

import (
	"github.com/stretchr/testify/require"
	"pkg/chess-logic/result"
	"testing"
)

func TestCell(t *testing.T) {
	testCases := []struct {
		notation string
		row      int
		col      int
	}{
		{"a1", 7, 0},
		{"h8", 0, 7},
		{"g4", 4, 6},
		{"c6", 2, 2},
	}

	p := NewStartPosition()

	t.Run("GetCell", func(t *testing.T) {
		for _, tc := range testCases {
			c, err := p.GetCell(tc.row, tc.col)
			require.NoError(t, err)
			require.Equal(t, tc.row, c.Row())
			require.Equal(t, tc.col, c.Col())
		}
	})

	t.Run("GetCellByNotation", func(t *testing.T) {
		for _, tc := range testCases {
			c, err := p.GetCellByNotation(tc.notation)
			require.NoError(t, err)
			require.Equal(t, tc.row, c.Row())
			require.Equal(t, tc.col, c.Col())
		}
	})

	t.Run("GetNotationByCell", func(t *testing.T) {
		for _, tc := range testCases {
			c, _ := p.GetCellByNotation(tc.notation)
			require.Equal(t, tc.notation, p.GetNotationByCell(c))
		}
	})
}

func TestGetCellByNotationError(t *testing.T) {
	testCases := []string{
		"h10", "x", "", "a",
		"4", "a9", "p7", "3f",
		"ab", "34", "a4a",
	}
	p := NewStartPosition()
	for _, tc := range testCases {
		_, err := p.GetCellByNotation(tc)
		require.Error(t, err, tc)
	}
}

func TestGame(t *testing.T) {
	testCases := []struct {
		startFen string
		moves    []string
		finalFen string
	}{

		{
			startFen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			moves:    []string{"e2e4", "e7e5", "d2d4", "d7d5", "e4d5", "e5d4", "d1d4", "d8e7"},
			finalFen: "rnb1kbnr/ppp1qppp/8/3P4/3Q4/8/PPP2PPP/RNB1KBNR w KQkq - 1 5",
		},

		{
			startFen: "rnbqkbnr/ppp2ppp/8/3p4/3QP3/8/PPP2PPP/RNB1KBNR w KQkq d6 0 4",
			moves:    []string{"e4e5", "f7f5", "e5f6", "g8f6", "g1f3", "f8d6", "h1g1", "e8g8", "c2c4"},
			finalFen: "rnbq1rk1/ppp3pp/3b1n2/3p4/2PQ4/5N2/PP3PPP/RNB1KBR1 b Q c3 0 8",
		},
		{
			startFen: "r3kb1r/ppp1q1pp/3pb3/5p2/2B1nB2/2N5/PPQ2PPP/3RR1K1 w kq - 4 12",
			moves:    []string{"c4e6", "e7e6", "c3e4", "f5e4", "c2e4", "e6e4", "e1e4", "e8d7", "d1e1", "h7h5"},
			finalFen: "r4b1r/pppk2p1/3p4/7p/4RB2/8/PP3PPP/4R1K1 w - h6 0 17",
		},
		{
			startFen: "r3k2r/p1ppbppp/bp6/8/8/PBN5/1P1B1PPP/R3K2R b KQkq - 2 12",
			moves:    []string{"a8d8", "h1f1", "e8g8", "h2h4", "e7h4", "b3a2", "a6b7", "a2b3", "b7a6"},
			finalFen: "3r1rk1/p1pp1ppp/bp6/8/7b/PBN5/1P1B1PP1/R3KR2 w Q - 4 17",
		},
		{
			startFen: "r1bqk2r/pppp1ppp/2n2n2/8/1b2P3/2N1Q3/PPPB1PPP/R3KBNR b KQkq - 6 6",
			moves: []string{"e8g8", "a2a3", "b4a5", "e1c1", "h7h6", "e4e5", "d7d5",
				"e5d6", "a5c3", "b2c3", "f6g4", "d6c7", "g4e3", "c7d8n"},
			finalFen: "r1bN1rk1/pp3pp1/2n4p/8/8/P1P1n3/2PB1PPP/2KR1BNR b - - 0 13",
		},
		{
			startFen: "8/P7/8/7k/8/7K/p7/1N6 w - - 0 1",
			moves:    []string{"a7a8r", "a2b1q", "a8h8", "b1h7"},
			finalFen: "7R/7q/8/7k/8/7K/8/8 w - - 2 3",
		},
	}

	for _, tc := range testCases {
		p, err := NewFromFEN(tc.startFen)
		require.NoError(t, err, tc.startFen)
		for _, move := range tc.moves {
			res, err := p.MakeMoveByString(move)
			require.NoError(t, err, tc.startFen, move)
			require.Equal(t, result.Unknown, res)
		}
		require.Equal(t, tc.finalFen, p.GetFEN())
	}
}

func TestGameResult(t *testing.T) {
	testCases := []struct {
		startFen string
		moves    []string
		finalFen string
		result   result.Result
	}{

		{
			startFen: "rnbqkbnr/ppppp2p/8/5pp1/4P3/2N5/PPPP1PPP/R1BQKBNR w KQkq g6 0 3",
			moves:    []string{"d1h5"},
			result:   result.Win,
			finalFen: "rnbqkbnr/ppppp2p/8/5ppQ/4P3/2N5/PPPP1PPP/R1B1KBNR b KQkq - 1 3",
		},

		{
			startFen: "6k1/8/8/r7/8/8/7r/4K3 b - - 0 1",
			moves:    []string{"a5a1"},
			result:   result.Win,
			finalFen: "6k1/8/8/8/8/8/7r/r3K3 w - - 1 2",
		},

		{
			startFen: "4rkr1/4p1p1/8/8/8/8/6PP/4K2R w K - 1 1",
			moves:    []string{"e1g1"},
			result:   result.Win,
			finalFen: "4rkr1/4p1p1/8/8/8/8/6PP/5RK1 b - - 2 1",
		},
		{
			startFen: "7k/R1P5/8/8/8/8/8/7K w - - 0 1",
			moves:    []string{"c7c8r"},
			result:   result.Win,
			finalFen: "2R4k/R7/8/8/8/8/8/7K b - - 0 1",
		},
		{
			startFen: "8/8/8/6Q1/4k3/8/3Q4/Q6K w - - 0 1",
			moves:    []string{"a1f1"},
			result:   result.Draw,
			finalFen: "8/8/8/6Q1/4k3/8/3Q4/5Q1K b - - 1 1",
		},
		{
			startFen: "5k2/8/3p4/8/2B2b2/3P4/8/5K2 b - - 7 4",
			moves:    []string{"f4e5", "c4b5", "e5f4", "b5c4", "f4e5", "c4b5", "e5f4", "b5c4"},
			result:   result.Draw,
			finalFen: "5k2/8/3p4/8/2B2b2/3P4/8/5K2 b - - 15 8",
		},
	}

	for _, tc := range testCases {
		p, err := NewFromFEN(tc.startFen)
		require.NoError(t, err, tc.startFen)
		var res result.Result
		for _, move := range tc.moves {
			res, err = p.MakeMoveByString(move)
			require.NoError(t, err, tc.startFen, move)
		}
		require.Equal(t, tc.result, res)
		require.Equal(t, tc.finalFen, p.GetFEN())
	}
}
