package position

import (
	"github.com/stretchr/testify/require"
	"pkg/chess-logic/pieces"
	"testing"
)

func TestFEN(t *testing.T) {
	testCases := []struct {
		fen            string
		position       string
		move           byte
		castling       string
		enPassantCell  string
		halfMoveClock  int
		fullMoveNumber int
	}{
		{
			fen:            "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			position:       "rnbqkbnr\npppppppp\n--------\n--------\n--------\n--------\nPPPPPPPP\nRNBQKBNR\n",
			move:           'w',
			castling:       "KQkq",
			enPassantCell:  "-",
			halfMoveClock:  0,
			fullMoveNumber: 1,
		},
		{
			fen:            "r1bqk2r/1pp1bppp/2np1n2/p3p3/4P3/2NPBN2/PPP1BPPP/R2QK2R w KQkq a6 0 7",
			position:       "r-bqk--r\n-pp-bppp\n--np-n--\np---p---\n----P---\n--NPBN--\nPPP-BPPP\nR--QK--R\n",
			move:           'w',
			castling:       "KQkq",
			enPassantCell:  "a6",
			halfMoveClock:  0,
			fullMoveNumber: 7,
		},
		{
			fen:            "r1bqk2r/1pp1bpp1/2np1n1p/p3p3/P3P3/2NPBN1P/1PP1BPP1/1R1Q1RK1 b k - 6 12",
			position:       "r-bqk--r\n-pp-bpp-\n--np-n-p\np---p---\nP---P---\n--NPBN-P\n-PP-BPP-\n-R-Q-RK-\n",
			move:           'b',
			castling:       "k",
			enPassantCell:  "-",
			halfMoveClock:  6,
			fullMoveNumber: 12,
		},
		{
			fen:            "1rbqk3/1pp2ppr/5n1p/p2pp1B1/P3P3/1PNP1P2/2PQK1PP/R4Bb1 w - - 8 17",
			position:       "-rbqk---\n-pp--ppr\n-----n-p\np--pp-B-\nP---P---\n-PNP-P--\n--PQK-PP\nR----Bb-\n",
			move:           'w',
			castling:       "-",
			enPassantCell:  "-",
			halfMoveClock:  8,
			fullMoveNumber: 17,
		},
		{
			fen:            "8/8/2Q5/4pb2/2pb4/3k4/8/4K3 w - - 10 81",
			position:       "--------\n--------\n--Q-----\n----pb--\n--pb----\n---k----\n--------\n----K---\n",
			move:           'w',
			castling:       "-",
			enPassantCell:  "-",
			halfMoveClock:  10,
			fullMoveNumber: 81,
		},
		{
			fen:            "1rbqkb1r/ppp2ppp/3p1n2/8/P1BQP3/8/1PP2PPP/RNB1K1R1 b Qk a3 0 8",
			position:       "-rbqkb-r\nppp--ppp\n---p-n--\n--------\nP-BQP---\n--------\n-PP--PPP\nRNB-K-R-\n",
			move:           'b',
			castling:       "Qk",
			enPassantCell:  "a3",
			halfMoveClock:  0,
			fullMoveNumber: 8,
		},
		{
			fen:            "rnbqkbnr/ppp1pppp/3p4/8/Q7/2P5/PP1PPPPP/RNB1KBNR b KQkq - 1 2",
			position:       "rnbqkbnr\nppp-pppp\n---p----\n--------\nQ-------\n--P-----\nPP-PPPPP\nRNB-KBNR\n",
			move:           'b',
			castling:       "KQkq",
			enPassantCell:  "-",
			halfMoveClock:  1,
			fullMoveNumber: 2,
		},
		{
			fen:            "rn3b1r/ppp1pppp/3p1n2/3k4/3P2b1/2P1PN2/PP1N1PPP/R1B1KBR1 w Q - 5 9",
			position:       "rn---b-r\nppp-pppp\n---p-n--\n---k----\n---P--b-\n--P-PN--\nPP-N-PPP\nR-B-KBR-\n",
			move:           'w',
			castling:       "Q",
			enPassantCell:  "-",
			halfMoveClock:  5,
			fullMoveNumber: 9,
		},
	}

	for _, tc := range testCases {
		p, err := NewFromFEN(tc.fen)
		require.NoError(t, err, tc.fen)
		require.Equal(t, tc.position, p.String())
		require.Equal(t, tc.move, pieces.GetByteByColor(p.whoseMove))
		require.Equal(t, tc.castling, p.castling.String())
		require.Equal(t, tc.enPassantCell, p.GetNotationByCell(p.enPassantCell))
		require.Equal(t, tc.halfMoveClock, p.halfMoveClock)
		require.Equal(t, tc.fullMoveNumber, p.fullMoveNumber)
		require.Equal(t, tc.fen, p.GetFEN())
	}
}

func TestFENError(t *testing.T) {
	testCases := []struct {
		fen string
	}{
		{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0"},
		{"rnbqkbnr/pppppppp/8/8/8/7/PPPPPPPP/RNBQKBNR w KQkq - 0 1"},
		{"rnbqkbnr/pppppppp/8/8/17/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"},
		{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR g KQkq - 0 1"},
		{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 -10"},
		{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - -8 1"},
		{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQqk - 1 10"},
		{"1rbqkb1r/ppp2ppp/3p1n2/8/P1BQP3/8/1PP2PPP/RNB1K1R1 b Qk a9 0 8"},
		{"1rbqkb1r/ppp2ppp/3p1n2/8/P1BQP3/8/1PP2PPP/RNB1K1R1 b Qk x4 0 8"},
		{"1rbqkb1r/ppp2ppp/3p1n2/8/P1BQP3/8/1PP2PPP/RNB1K1R1 b Qk a33 0 8"},
		{"1rbqkb1r/ppp2ppp/3p1n2/8/P1BQP3/8/1PP2PPP/RNB1K1R1 b kQ a3 0 8"},
		{"1rbqkb1r/ppp11ppp/3p1n2/8/P1BQP3/8/1PP2PPP/RNB1K1R1 b Qk a3 0 8"},
		{"1rbqkb1r/ppp2ppp/3p1n2/8/P1BQP3/8/1P2PPP/RNB1K1R1 b Qk a3 0 8"},
		{"1rbqkb1r/ppp2ppp/3p1n2/8/P1BQP3/8/1PP2PPP/ b Qk a3 0 8"},
		{"1rbqkb1r/ppp2ppp/3p1n2/8/P1BQP3/8/1PP2PPP b Qk a3 0 8"},
		{"1rbqkb1r/ppp2ppp/3p1n2/8/P1BQP3/8/1PP2PPP/RNB1K1R1 bb Qk a3 0 8"},
		{"1rbq1b1r/ppp2ppp/3p1n2/8/P1BQP3/8/1PP2PPP/RNB1K1R1 b Qk a3 0 8"},
		{"1rbqkb1r/ppp2ppp/3p1n2/8/P1BQP3/8/1PP2PPP/RNB3R1 b Qk a3 0 8"},
		{"1rb1kb1r/ppp1qppp/3p4/6n1/P1BQ4/2N2P2/1PP3PP/R1B1K1R1 b Qk - 1 11"},
		{"rnbqkbnr/pppdpppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"},
		{"1rb1kb1r/ppp1qppp/3p4/6n1/P1BQ4/2N2P2/1PP3PP/K1B1K1R1 w Qk - 1 11"},
		{"1rb1kb1r/ppp1qppp/3p4/6n2/P1BQ4/2N2P2/1PP3PP/R1B1K1R1 w Qk - 1 11"},
		{"rnbqkbnr/pppplppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"},
		{"rnb1kbnr/pppp1ppp/4p3/8/5P1q/2N5/PPPPP1PP/R1BQKBNR b KQkq - 2 3"},
		{"rn3b1r/ppp1pppp/3p1n2/3k4/2PP2b1/4PN2/PP1N1PPP/R1B1KBR1 w Q - 0 9"},
		{"rnbqkbnr/ppp1pppp/3p4/8/Q7/2P5/PP1PPPPP/RNB1KBNR w KQkq - 1 2"},
		{"8/8/8/8/8/8/3Kk3/8 w - - 0 1"},
	}

	for _, tc := range testCases {
		_, err := NewFromFEN(tc.fen)
		require.Error(t, err, tc.fen)
	}
}
