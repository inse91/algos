package _1812_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/determine-color-of-a-chessboard-square/description/

func Test(t *testing.T) {
	cases := []struct {
		name   string
		coords string
		res    bool
	}{
		{
			name:   "1",
			coords: "a1",
			res:    false,
		},
		{
			name:   "2",
			coords: "h3",
			res:    true,
		},
		{
			name:   "3",
			coords: "c7",
			res:    false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, squareIsWhite(c.coords))
		})
	}
}

func squareIsWhite(coordinates string) bool {
	if len(coordinates) < 2 {
		panic("invalid coordinates")
	}

	letter := coordinates[0]
	digit := coordinates[1]

	if letter%2 == 0 {
		return digit%2 == 1
	}

	return digit%2 == 0
}
