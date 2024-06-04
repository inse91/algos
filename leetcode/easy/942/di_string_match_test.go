package _942_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/di-string-match/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  []int
	}{
		{
			name: "1",
			s:    "IDID",
			res:  []int{0, 4, 1, 3, 2},
		},
		{
			name: "2",
			s:    "III",
			res:  []int{0, 1, 2, 3},
		},
		{
			name: "3",
			s:    "DDI",
			res:  []int{3, 2, 0, 1},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, diStringMatch(c.s))
		})
	}
}

func diStringMatch(s string) []int {
	res := make([]int, 0, len(s)+1)

	start := 0
	end := len(s)
	for _, r := range s {
		switch r {
		case 'D':
			res = append(res, end)
			end--
		case 'I':
			res = append(res, start)
			start++
		default:
			panic("invalid letter")
		}
	}

	res = append(res, end)

	return res
}
