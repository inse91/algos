package _2194_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/cells-in-a-range-on-an-excel-sheet

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  []string
	}{
		{
			name: "1",
			s:    "K1:L2",
			res:  []string{"K1", "K2", "L1", "L2"},
		},
		{
			name: "1",
			s:    "A1:F1",
			res:  []string{"A1", "B1", "C1", "D1", "E1", "F1"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, cellsInRange(c.s))
		})
	}
}

func cellsInRange(s string) []string {
	codes := strings.Split(s, ":")
	if len(codes) != 2 || len(codes[0]) != 2 || len(codes[1]) != 2 {
		return nil
	}

	start := codes[0]
	end := codes[1]
	startLetter := start[0]
	endLetter := end[0]
	startDig := start[1]
	endDig := end[1]

	res := make([]string, 0, (endLetter-startLetter+1)*(endDig-startDig+1))

	var sb strings.Builder
	for i := startLetter; i <= endLetter; i++ {
		for j := startDig; j <= endDig; j++ {
			sb.WriteByte(i)
			sb.WriteByte(j)

			res = append(res, sb.String())

			sb.Reset()
		}
	}

	return res
}
