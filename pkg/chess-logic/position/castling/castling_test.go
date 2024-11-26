package castling

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewByString(t *testing.T) {
	testCases := []struct {
		str  string
		want *castlingImpl
	}{
		{"KQkq", &castlingImpl{
			whiteKingSideCastling:  true,
			whiteQueenSideCastling: true,
			blackKingSideCastling:  true,
			blackQueenSideCastling: true,
		}},
		{"Kq", &castlingImpl{
			whiteKingSideCastling:  true,
			whiteQueenSideCastling: false,
			blackKingSideCastling:  false,
			blackQueenSideCastling: true,
		}},
		{"K", &castlingImpl{
			whiteKingSideCastling:  true,
			whiteQueenSideCastling: false,
			blackKingSideCastling:  false,
			blackQueenSideCastling: false,
		}},
		{"q", &castlingImpl{
			whiteKingSideCastling:  false,
			whiteQueenSideCastling: false,
			blackKingSideCastling:  false,
			blackQueenSideCastling: true,
		}},
		{"-", &castlingImpl{
			whiteKingSideCastling:  false,
			whiteQueenSideCastling: false,
			blackKingSideCastling:  false,
			blackQueenSideCastling: false,
		}},
	}
	for _, tc := range testCases {
		c, err := NewByString(tc.str)
		require.Equal(t, tc.want, c)
		require.NoError(t, err)
		require.Equal(t, tc.str, c.String())
	}
}

func TestByStringError(t *testing.T) {
	testCases := []string{
		"", "KKQkq", "L", "KKq",
		"kK", "QK", "qk", "Kql",
		"Kqq", "KQK", "kqK", "KQqk",
	}

	for _, tc := range testCases {
		_, err := NewByString(tc)
		require.Error(t, err)
	}
}
