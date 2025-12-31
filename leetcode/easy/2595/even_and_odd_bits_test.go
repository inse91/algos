package _2595_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-even-and-odd-bits/description

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  []int
	}{
		{
			name: "1",
			n:    50,
			res:  []int{1, 2},
		},
		{
			name: "2",
			n:    2,
			res:  []int{0, 1},
		},
		{
			name: "737",
			n:    5,
			res:  []int{2, 0},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, evenOddBit(c.n))
		})
	}
}

func evenOddBit(n int) []int {
	var (
		odds, evens int
	)
	beets := strconv.FormatInt(int64(n), 2)
	for i := range beets {
		b := beets[len(beets)-i-1]
		if b != '1' {
			continue
		}

		switch i % 2 {
		case 0:
			evens++
		case 1:
			odds++
		}
	}

	return []int{evens, odds}
}
