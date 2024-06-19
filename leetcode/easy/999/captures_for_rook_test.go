package _999_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/available-captures-for-rook/description/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		board [][]byte
		res   int
	}{
		{
			name: "1",
			board: [][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', 'R', '.', '.', '.', 'p'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
			},
			res: 3,
		},
		{
			name: "2",
			board: [][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
				{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
				{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
				{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
				{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
			},
			res: 0,
		},
		{
			name: "3",
			board: [][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', 'p'},
				{'p', 'p', '.', 'R', '.', 'p', 'B', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'B', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
			},
			res: 3,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numRookCaptures(c.board))
		})
	}
}

const (
	rook byte = 'R'
	bish byte = 'B'
	pawn byte = 'p'
)

func numRookCaptures(board [][]byte) int {
	rookPoint, found := findRook(board)
	if !found {
		return 0
	}

	var c int
	x, y := rookPoint[0], rookPoint[1]

	lineX := board[x]
	lineY := make([]byte, len(board))
	for i := range board {
		lineY[i] = board[i][y]
	}

	// ==>> x
	for i := y + 1; i < len(lineX); i++ {
		if lineX[i] == bish {
			break
		}
		if lineX[i] == pawn {
			c++
			break
		}
	}

	// <<== x
	for i := y - 1; i > -1; i-- {
		if lineX[i] == bish {
			break
		}
		if lineX[i] == pawn {
			c++
			break
		}
	}

	// ==>> y
	for i := x + 1; i < len(lineY); i++ {
		if lineY[i] == bish {
			break
		}
		if lineY[i] == pawn {
			c++
			break
		}
	}

	// ==>> y
	for i := x - 1; i > -1; i-- {
		if lineY[i] == bish {
			break
		}
		if lineY[i] == pawn {
			c++
			break
		}
	}

	return c
}

func findRook(board [][]byte) ([2]int, bool) {
	for i, bts := range board {
		for j, b := range bts {
			if b == rook {
				return [2]int{i, j}, true
			}
		}
	}

	return [2]int{}, false
}
